package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytesToStruct(t *testing.T) {
	t.Run("returns error when too few bytes given", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
		}
		err := BytesToStruct(bytes, &Vec3f{}, 12)

		expectedErr := NewByteCountMismatchError("*common.Vec3f", 12, 8)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("returns error when too many bytes given", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
			0x00, 0x00, 0x40, 0x40, // 3.0
			0x00, 0x00, 0x80, 0x3f, // 1.0
		}
		err := BytesToStruct(bytes, &Vec3f{}, 12)

		expectedErr := NewByteCountMismatchError("*common.Vec3f", 12, 16)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("correctly decodes struct pointer", func(t *testing.T) {
		bytes := []byte{
			0x00, 0x00, 0x80, 0x3f, // 1.0
			0x00, 0x00, 0x00, 0x40, // 2.0
			0x00, 0x00, 0x40, 0x40, // 3.0
		}
		actual := &Vec3f{}
		err := BytesToStruct(bytes, actual, 12)

		expected := NewVec3f(1.0, 2.0, 3.0)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}
