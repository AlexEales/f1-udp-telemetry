package v2019

import "f1-udp-telemetry/data/common"

// TODO: Need to add in some of the enums required for this

// SessionPacketByteCount represents the number of bytes that make up a session packet
const SessionPacketByteCount = 149

// SessionPacket contains all the information about the session taking place
type SessionPacket struct {
	Header                    PacketHeader
	Weather                   common.Weather
	TrackTemperatureInCelsius int8
	AirTemperatureInCelsius   int8
	TotalLaps                 uint8
	TrackLengthInMeters       uint16
	SessionType               common.SessionType
	TrackID                   int8
	Formula                   uint8
	SessionTimeLeftInSeconds  uint16
	SessionDurationInSeconds  uint16
	PitSpeedLimitInKPH        uint8
	IsGamePaused              bool
	AreSpectating             bool
	SpectatingCarIndex        uint8
	SLIProNativeIsSupported   bool
	NumMarshalZones           uint8
	MarshalZones              [21]MarshalZone
	SafetyCarStatus           uint8
	IsNetworkGame             bool
}

// MarshalZone represents a area of the track which has a marshal stationed and what the status of that zone is
type MarshalZone struct {
	// ZoneStart is a fraction 0-1 representing how far through the lap the marshal zone starts
	ZoneStart float32
	// ZoneFlag indicates the flag being flown in that marshal zone
	ZoneFlag common.ZoneFlag
}
