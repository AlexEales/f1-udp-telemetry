package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

// BytesToStruct takes in bytes, pointer to target struct and the expected byte count and attempts to
// read the bytes into the target struct.
func BytesToStruct(data []byte, target interface{}, targetByteCount int) error {
	if len(data) != targetByteCount {
		typeStr := fmt.Sprintf("%T", target)
		return NewByteCountMismatchError(typeStr, targetByteCount, len(data))
	}

	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, target)
	return nil
}
