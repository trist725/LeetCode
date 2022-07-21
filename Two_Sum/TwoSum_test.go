package Two_Sum

import "testing"

func BenchmarkTwoSum_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum_1([]int{2, 11, 15, 5, 6, 8, 9, 2, 3, 5, 7, 9, 0, 9, 454, 6, 578, 97, 07, 43}, 14)
	}
}

func BenchmarkTwoSum_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum_2([]int{2, 11, 15, 5, 6, 8, 9, 2, 3, 5, 7, 9, 0, 9, 454, 6, 578, 97, 07, 43}, 14)
	}
}
