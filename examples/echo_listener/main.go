package main

import (
	"context"
	"f1-udp-telemetry/data/common"
	"f1-udp-telemetry/listener"

	log "github.com/sirupsen/logrus"
)

func handleCarStatusPacket(bytes []byte) {
	log.WithField("fuel mix", bytes[26]).Info("Received car status packet")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l := listener.NewPacketListener()
	l.AddHandler(common.CarStatusPacket, handleCarStatusPacket)

	err := l.Start(ctx)
	if err != nil {
		log.WithError(err).Error("error occured while listening for packets")
	}
}
