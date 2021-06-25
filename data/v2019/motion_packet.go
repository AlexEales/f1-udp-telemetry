package v2019

import "f1-udp-telemetry/data/common"

// MotionPacketByteCount represents the number of bytes that make up a motion packet
const MotionPacketByteCount = 1343

// MotionPacket represnts a F12019 motion packet containing information about the physiscs data of all cars
type MotionPacket struct {
	Header     PacketHeader
	MotionData [20]CarMotionData

	// Player car only data

	// SuspensionPosition is the position of each wheels suspension (order RL, RR, FL, FR)
	SuspensionPosition common.Vec4f
	// SuspensionVelocity is the velocity of each wheels suspension (order RL, RR, FL, FR)
	SuspensionVelocity common.Vec4f
	// SuspensionAcceleration is the acceleration of each wheels suspension (order RL, RR, FL, FR)
	SuspensionAcceleration common.Vec4f
	// WheelSpin is the spin of each of the wheels (order RL, RR, FL, FR)
	WheelSpin common.Vec4f
	// Wheel slip is the slip of each of the wheels (order RL, RR, FL, FR)
	WheelSlip           common.Vec4f
	LocalVelocity       common.Vec3f
	AngularVelocity     common.Vec3f
	AngularAcceleration common.Vec3f
	FrontWheelsAngle    float32
}

// CarMotionData contains all the information about the motion data of a particular car
type CarMotionData struct {
	WorldPosition common.Vec3f
	Velocity      common.Vec3f
	// WorldForwardDirection is the normalised world space forward direction
	WorldForwardDirection common.Vec3i16
	// WorldRightDirection is the normalised world space right direction
	WorldRightDirection common.Vec3i16
	// GForce is the G-Forces acting on the car in lateral, longitudinal and vertical directions (X, Y and Z)
	GForce common.Vec3f
	// Rotation is the yaw, pitch and roll (X, Y and Z) of the car in radians
	Rotation common.Vec3f
}
