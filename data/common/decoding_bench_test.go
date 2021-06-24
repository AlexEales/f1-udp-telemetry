package common

import (
	"math/rand"
	"testing"
)

func BenchmarkBytesToFloat(b *testing.B) {
	testSet := generateTestSet(b.N, 4)
	for n := 0; n < b.N; n++ {
		if _, err := BytesToFloat(testSet[n]); err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkBytesToVec3f(b *testing.B) {
	testSet := generateTestSet(b.N, 12)
	for n := 0; n < b.N; n++ {
		if _, err := BytesToVec3f(testSet[n]); err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkBytesToVec4f(b *testing.B) {
	testSet := generateTestSet(b.N, 16)
	for n := 0; n < b.N; n++ {
		if _, err := BytesToVec4f(testSet[n]); err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkBytesToVec3u16(b *testing.B) {
	testSet := generateTestSet(b.N, 6)
	for n := 0; n < b.N; n++ {
		if _, err := BytesToVec3u16(testSet[n]); err != nil {
			b.FailNow()
		}
	}
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
