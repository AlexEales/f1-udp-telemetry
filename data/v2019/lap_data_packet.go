package v2019

import "f1-udp-telemetry/data/common"

// LapDataPacketByteCount represents the number of bytes that make up a lap data packet
const LapDataPacketByteCount = 843

// LapDataPacket contains all the information about the lap data of the cars in the session
type LapDataPacket struct {
	Header PacketHeader
	Data   [20]LapData
}

// LapData contains the lap data for a individual car
type LapData struct {
	LastLapTimeInSeconds       float32
	CurrentLapTimeInSeconds    float32
	BestLapTimeInSeconds       float32
	Sector1TimeInSeconds       float32
	Sector2TimeInSeconds       float32
	CurrentLapDistanceInMeters float32
	TotalLapDistanceInMeters   float32
	SafetyCarDeltaInSeconds    float32
	CarPosition                uint8
	CurrentLapNumber           uint8
	PitStatus                  common.PitStatus
	Sector                     common.Sector
	IsCurrentLapInvalid        bool
	PenaltiesInSeconds         uint8
	GridPosition               uint8
	DriverStatus               common.DriverStatus
	ResultStatus               common.ResultStatus
}
