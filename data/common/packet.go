package common

import "fmt"

type Packet interface {
	Name() string
	Version() PacketFormat
}

// PacketFormat defines the format of the packet based on the game version
type PacketFormat uint16

var (
	// F1_2019 identifies the packet format for the F1 2019 game
	F1_2019 PacketFormat = 2019
	// F1_2029 identifies the packet format for the F1 2020 game
	F1_2020 PacketFormat = 2020
)

// PacketFormat_String map from packet format to it's string name
var PacketFormat_String = map[PacketFormat]string{
	F1_2019: "F1 2019",
	F1_2020: "F1 2020",
}

// String returns the string version of the packet format with
// "format" appended to the end
func (f PacketFormat) String() string {
	return fmt.Sprintf("%s format", PacketFormat_String[f])
}

// Valid returns if the packet format is considered valid (has a entry in the string
// dictionary)
func (f PacketFormat) Valid() bool {
	_, ok := PacketFormat_String[f]
	return ok
}

// PacketType defines the type of the packet
type PacketType uint8

var (
	// MotionPacket identifies a Motion Data Packet
	MotionPacket PacketType = 0
	// SessionPacket identifies a Session Data Packet
	SessionPacket PacketType = 1
	// LapDataPacket identifies a Lap Data Packet
	LapDataPacket PacketType = 2
	// EventPacket identifies a Event Data packet
	EventPacket PacketType = 3
	// ParticipantsPacket identifies a Participants Data Packet
	ParticipantsPacket PacketType = 4
	// CarSetupsPacket identifies a Car Setups Data Packet
	CarSetupsPacket PacketType = 5
	// CarTelemetryPacket identifies a Car Telemetry Data Packet
	CarTelemetryPacket PacketType = 6
	// CarStatusPacket identifies a Car Status Data Packet
	CarStatusPacket PacketType = 7
	// FinalClassificationPacket identifies a Final Classification Data Packet
	FinalClassificationPacket PacketType = 8
	// LobbyInfoPacket identifies a Lobby Info Packet
	LobbyInfoPacket PacketType = 9
	// AnyPacket is a wildcard packet identifier
	AnyPacket PacketType = 254
)

// PacketType_String map from packet type to it's string name
var PacketType_String = map[PacketType]string{
	MotionPacket:              "Motion Packet",
	SessionPacket:             "Session Packet",
	LapDataPacket:             "Lap Data Packet",
	EventPacket:               "Event Packet",
	ParticipantsPacket:        "Participants Packet",
	CarSetupsPacket:           "Car Setups Packet",
	CarTelemetryPacket:        "Car Telemetry Packet",
	CarStatusPacket:           "Car Status Packet",
	FinalClassificationPacket: "Final Classification Packet",
	LobbyInfoPacket:           "Lobby Info Packet",
	AnyPacket:                 "Any Packet",
}

// String returns the string value for the packet type
func (t PacketType) String() string {
	return PacketType_String[t]
}

// Valid returns if the packet type is considered valid (has a entry in the string
// dictionary)
func (t PacketType) Valid() bool {
	_, ok := PacketType_String[t]
	return ok
}
