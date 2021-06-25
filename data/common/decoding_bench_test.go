package common

import (
	"math/rand"
	"testing"
)

// TODO: Move this into common decoding folder since we would want to test 2019 and 2020 packets in here too?

func BenchmarkBytesToStruct(b *testing.B) {
	b.Run("vec3f", func(b *testing.B) {
		testSet := generateTestSet(b.N, 12)
		for n := 0; n < b.N; n++ {
			if err := BytesToStruct(testSet[n], &Vec3f{}, 12); err != nil {
				b.FailNow()
			}
		}
	})

	b.Run("vec4f", func(b *testing.B) {
		testSet := generateTestSet(b.N, 16)
		for n := 0; n < b.N; n++ {
			if err := BytesToStruct(testSet[n], &Vec4f{}, 16); err != nil {
				b.FailNow()
			}
		}
	})

	b.Run("vec3i16", func(b *testing.B) {
		testSet := generateTestSet(b.N, 6)
		for n := 0; n < b.N; n++ {
			if err := BytesToStruct(testSet[n], &Vec3i16{}, 6); err != nil {
				b.FailNow()
			}
		}
	})
}

func generateTestSet(tests, size int) [][]byte {
	set := [][]byte{}
	for t := 0; t < tests; t++ {
		bytes := make([]byte, size)
		rand.Read(bytes)
		set = append(set, bytes)
	}
	return set
}
