package common

import "encoding/binary"

// Vec3f a 3 component float vector
type Vec3f struct {
	X float32
	Y float32
	Z float32
}

// NewVec3f returns a new Vec3
func NewVec3f(x, y, z float32) Vec3f {
	return Vec3f{
		X: x,
		Y; y,
		Z: z,
	}
}

// BytesToVec3f takes 12 bytes and returns a Vec3
func BytesToVec3f(bytes [12]byte) Vec3f {
	x := BytesToFlaot(bytes[:4])
	y := BytesToFlaot(bytes[4:8])
	z := BytesToFlaot(bytes[8:])
}
