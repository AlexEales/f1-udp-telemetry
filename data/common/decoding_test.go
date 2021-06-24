package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytesToFloat(t *testing.T) {
	t.Run("returns error when too few bytes given", func(t *testing.T) {
		bytes := []byte{0x56, 0x0e}
		_, err := BytesToFloat(bytes)

		expectedErr := NewByteCountMismatchError("float", 4, 2)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns error when too many bytes given", func(t *testing.T) {
		bytes := []byte{0x56, 0x0e, 0x01, 0x02, 0x03, 0x04}
		_, err := BytesToFloat(bytes)

		expectedErr := NewByteCountMismatchError("float", 4, 6)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns correct float", func(t *testing.T) {
		bytes := []byte{0x56, 0x0e, 0x49, 0x40}
		actual, err := BytesToFloat(bytes)

		expected := float32(3.1415)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}

func TestBytesToVec3f(t *testing.T) {
	t.Run("returns error when too few bytes given", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
		}
		_, err := BytesToVec3f(bytes)

		expectedErr := NewByteCountMismatchError("Vec3f", 12, 8)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns error when too many bytes given", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
			0x00, 0x00, 0x40, 0x40, // 3.0
			0x00, 0x00, 0x80, 0x3f, // 1.0
		}
		_, err := BytesToVec3f(bytes)

		expectedErr := NewByteCountMismatchError("Vec3f", 12, 16)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns correct Vec3f", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
			0x00, 0x00, 0x40, 0x40, // 3.0
		}
		actual, err := BytesToVec3f(bytes)

		expected := NewVec3f(1.0, 2.0, 3.0)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}

func TestBytesToVec4f(t *testing.T) {
	t.Run("returns error when too few bytes given", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
			0x00, 0x00, 0x40, 0x40, // 3.0
		}
		_, err := BytesToVec4f(bytes)

		expectedErr := NewByteCountMismatchError("Vec4f", 16, 12)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns error when too many bytes given", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
			0x00, 0x00, 0x40, 0x40, // 3.0
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
		}
		_, err := BytesToVec4f(bytes)

		expectedErr := NewByteCountMismatchError("Vec4f", 16, 20)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns correct Vec4f", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
			0x00, 0x00, 0x40, 0x40, // 3.0
			0x00, 0x00, 0x80, 0x40, // 4.0
		}
		actual, err := BytesToVec4f(bytes)

		expected := NewVec4f(1.0, 2.0, 3.0, 4.0)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}

func TestBytesToVec3u16(t *testing.T) {
	t.Run("returns error when too few bytes given", func(t *testing.T) {
		bytes := []byte{
			0x01, 0x00, // 1
			0x02, 0x00, // 2
		}
		_, err := BytesToVec3u16(bytes)

		expectedErr := NewByteCountMismatchError("Vec3u16", 6, 4)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns error when too many bytes given", func(t *testing.T) {
		bytes := []byte{
			0x01, 0x00, // 1
			0x02, 0x00, // 2
			0x03, 0x00, // 3
			0x01, 0x00, // 1
		}
		_, err := BytesToVec3u16(bytes)

		expectedErr := NewByteCountMismatchError("Vec3u16", 6, 8)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns correct Vec3u16", func(t *testing.T) {
		bytes := []byte{
			0x01, 0x00, // 1
			0x02, 0x00, // 2
			0x03, 0x00, // 3
		}
		actual, err := BytesToVec3u16(bytes)

		expected := NewVec3u16(1, 2, 3)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}
