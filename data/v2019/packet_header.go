package v2019

import "f1-udp-telemetry/data/common"

// PacketHeaderByteCount represents the number of bytes that makes up a PacketHeader
const PacketHeaderByteCount = 24

// PacketHeader is the common header for all packets and contains information about the packet
type PacketHeader struct {
	// Format is the packet format (2020, 2019, ...)
	Format common.PacketFormat
	// MajorVersion is the major version of the game client sending the packet
	MajorVersion uint8
	// MinorVersion is the minor version of the game client sending the packet
	MinorVersion uint8
	// Version is the version of the packet being sent (all start from 1)
	Version uint8
	// Type is the type of the packet (Session, Motion, etc.)
	Type common.PacketType
	// SessionUID is a unique identifier for the session
	SessionUID uint64
	// SessionTime is the current timestamp in the session when the packet was sent
	SessionTime float32
	// FrameIdentifier identifies the frame the data was retrieved on
	FrameIdentifier uint
	// PlayerCarIndex is the index of the players car in player indexed arrays
	PlayerCarIndex uint8
}
