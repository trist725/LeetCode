package Remove_Element

import "testing"

func BenchmarkRemoveElement_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElement_1([]int{3, 2, 2, 3}, 3)
	}
}

func BenchmarkRemoveElement_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElement_2([]int{3, 2, 2, 3}, 3)
	}
}
