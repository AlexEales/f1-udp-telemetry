package f12019

import "f1-udp-telemetry/data/common"

// PacketHeader represents all the data contained in a F1 2019
// packet header
type PacketHeader struct {
	Format common.PacketFormat
	VersionMajor uint8
	VersionMinor uint8
	Version uint8
	Type common.PacketType
	SessionUID uint64
	SessionTime float32
	FrameIdentifier uint32
	PlayerCarIndex uint8
}
