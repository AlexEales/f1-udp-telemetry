package common

// Vec3f represents a 3-component float vector
type Vec3f struct {
	X, Y, Z float32
}

// NewVec3f returns a new Vec3f from 3 floats
func NewVec3f(x, y, z float32) *Vec3f {
	return &Vec3f{
		X: x,
		Y: y,
		Z: z,
	}
}

// Vec4f represents a 4-component float vector
type Vec4f struct {
	X, Y, Z, W float32
}

// NewVec4f returns a new Vec4f from 4 floats
func NewVec4f(x, y, z, w float32) *Vec4f {
	return &Vec4f{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// Vec3u16 represents a 3-component u16 vector
type Vec3u16 struct {
	X, Y, Z uint16
}

// NewVec3u16 returns a new Vec3u16 from 3 uint16's
func NewVec3u16(x, y, z uint16) *Vec3u16 {
	return &Vec3u16{
		X: x,
		Y: y,
		Z: z,
	}
}
