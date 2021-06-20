package common

import (
	"encoding/binary"
	"math"
)

// BytesToFlaot takes 4 bytes and returns a float32
func BytesToFlaot(bytes [4]byte) float32 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}