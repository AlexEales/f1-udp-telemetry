package listener

import (
	"context"
	"encoding/binary"
	"f1-udp-telemetry/data/common"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultListenerHost the default host the listener will listen for packets on
	DefaultListenerHost string = "127.0.0.1"
	// DefaultListenerPort the default port the listener will listen for packets on
	DefaultListenerPort uint16 = 20777
)

// PacketHandler function definition for packet handlers
type PacketHandler func([]byte)

// PacketListener interface for a packet listener
type PacketListener interface {
	AddHandler(packetType common.PacketType, handler PacketHandler)
	GetHandlers(packetType common.PacketType) []PacketHandler
	HandlePacket(packet []byte) error
	Start(ctx context.Context) error
}

type packetListener struct {
	addr     *net.UDPAddr
	handlers map[common.PacketType][]PacketHandler
}

// TODO: Make the host etc. configurable
// NewPacketListener returns a new packet listener instance
func NewPacketListener() PacketListener {
	return &packetListener{
		addr: &net.UDPAddr{
			IP:   net.ParseIP(DefaultListenerHost),
			Port: int(DefaultListenerPort),
		},
		handlers: make(map[common.PacketType][]PacketHandler),
	}
}

// AddHandler adds a handler for a given packet type (can be any) to the list of handlers
// for that type
func (l *packetListener) AddHandler(packetType common.PacketType, handler PacketHandler) {
	existingHandlers := l.handlers[packetType]
	l.handlers[packetType] = append(existingHandlers, handler)
}

// GetHandlers returns the handlers present for a given packet type
func (l *packetListener) GetHandlers(packetType common.PacketType) []PacketHandler {
	return l.handlers[packetType]
}

func (l *packetListener) HandlePacket(packet []byte) error {
	packetFormat := common.PacketFormat(
		binary.LittleEndian.Uint16(packet[0:2]),
	)
	if !packetFormat.Valid() {
		return fmt.Errorf("packet format not valid %v", packetFormat)
	}

	packetType := common.PacketType(packet[5])
	if !packetType.Valid() {
		return fmt.Errorf("packet type not valid %v", packetType)
	}

	log.WithFields(log.Fields{
		"format": packetFormat,
		"type":   packetType,
		"size":   len(packet),
	}).Info("received packet")

	handlers := append(l.handlers[common.AnyPacket], l.handlers[packetType]...)
	for _, handler := range handlers {
		handler(packet)
	}

	return nil
}

// Start begins listening for packets and will only exit on error
// reading packet or on context cancellation
func (l *packetListener) Start(ctx context.Context) error {
	conn, err := net.ListenUDP("udp", l.addr)
	if err != nil {
		err := fmt.Errorf("failed to listen for UDP packets: %w", err)
		log.WithFields(log.Fields{
			"IP":   l.addr.IP,
			"port": l.addr.Port,
		}).Error(err)
		return err
	}
	defer conn.Close()

	log.WithFields(log.Fields{
		"IP":   l.addr.IP,
		"port": l.addr.Port,
	}).Info("Listening for UDP packets")

	var buf [2048]byte
	for {
		select {
		case <-ctx.Done():
			log.Info("listener context cancelled, exiting gracefully")
			return nil
		default:
			rlen, _, err := conn.ReadFromUDP(buf[:])
			if err != nil {
				return err
			}

			if len(buf) < 6 {
				log.WithField("size", len(buf)).
					Warn("received too small buffer could not determine type")
				continue
			}

			if err = l.HandlePacket(buf[:rlen]); err != nil {
				log.WithError(err).Warn("error handling packet")
			}
		}
	}

	return nil
}
