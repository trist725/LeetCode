package Binary_Search

import (
	"testing"
)

func BenchmarkBinarySearch_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch_1([]int{-1, 0, 3, 5, 9, 12}, 9)
	}
}

func BenchmarkBinarySearch_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch_2([]int{-1, 0, 3, 5, 9, 12}, 9)
	}
}
