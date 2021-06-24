package common

import (
	"encoding/binary"
	"fmt"
	"math"
)

// DecodingErrorType identifies the type of error which has occurred when decoding
type DecodingErrorType uint8

const (
	// ByteCountMistmatch occurs when the passed byte array is not the expected length required to decode
	ByteCountMismatch = DecodingErrorType(0)
)

// DecodingError is the error type for decoding related errors
type DecodingError struct {
	DecodingErrorType
	Err error
}

// Error implements go's error interface
func (e *DecodingError) Error() string {
	return e.Err.Error()
}

// Unwrap implements go's error interface
func (e *DecodingError) Unwrap() error {
	return e.Err
}

// NewByteCountMismatchError takes an error and wraps it into a DecodingError with type ByteCountMismatch
func NewByteCountMismatchError(typeStr string, expected, actual int) *DecodingError {
	err := fmt.Errorf("failed to decode %s, expected %d bytes, got %d", typeStr, expected, actual)
	return &DecodingError{
		DecodingErrorType: ByteCountMismatch,
		Err:               err,
	}
}

// BytesToFloat takes a slice of 4 bytes and returns a decoded float or errors
func BytesToFloat(bytes []byte) (float32, error) {
	if len(bytes) != 4 {
		return 0.0, NewByteCountMismatchError("float", 4, len(bytes))
	}

	return bytesToFloat(bytes), nil
}

// BytesToVec3f takes a slice of 12 bytes and returns a decoded Vec3f or errors
func BytesToVec3f(bytes []byte) (*Vec3f, error) {
	if len(bytes) != 12 {
		return nil, NewByteCountMismatchError("Vec3f", 12, len(bytes))
	}

	return &Vec3f{
		X: bytesToFloat(bytes[:4]),
		Y: bytesToFloat(bytes[4:8]),
		Z: bytesToFloat(bytes[8:]),
	}, nil
}

// BytesToVec4f takes a slice of 16 bytes and returns a decoded Vec4f or errors
func BytesToVec4f(bytes []byte) (*Vec4f, error) {
	if len(bytes) != 16 {
		return nil, NewByteCountMismatchError("Vec4f", 16, len(bytes))
	}

	return &Vec4f{
		X: bytesToFloat(bytes[:4]),
		Y: bytesToFloat(bytes[4:8]),
		Z: bytesToFloat(bytes[8:12]),
		W: bytesToFloat(bytes[12:]),
	}, nil
}

// BytesToVec3u16 takes a slice of 6 bytes and returns a decoded Vec3u16 or errors
func BytesToVec3u16(bytes []byte) (*Vec3u16, error) {
	if len(bytes) != 6 {
		return nil, NewByteCountMismatchError("Vec3u16", 6, len(bytes))
	}

	return &Vec3u16{
		X: binary.LittleEndian.Uint16(bytes[:2]),
		Y: binary.LittleEndian.Uint16(bytes[2:4]),
		Z: binary.LittleEndian.Uint16(bytes[4:]),
	}, nil
}

// bytesToFloat is a internal conversion which is unchecked
func bytesToFloat(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes[:])
	return math.Float32frombits(bits)
}
