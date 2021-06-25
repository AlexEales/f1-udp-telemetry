package v2019

import "f1-udp-telemetry/data/common"

// MotionPacketByteCount represents the number of bytes that make up a motion packet
const MotionPacketByteCount = 1343

// MotionPacket represnts a F12019 motion packet containing information about the physiscs data of all cars
type MotionPacket struct {
	WorldPosition        common.Vec3f
	Velocity             common.Vec3f
	WorldForwardPosition common.Vec3i16
}
