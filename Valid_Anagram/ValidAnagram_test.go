package Valid_Anagram

import "testing"

func BenchmarkIsAnagram_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsAnagram_1("anagramanagramanagramanagramanagram", "nagaramnagaramnagaramnagaramnagaramnagaram")
	}
}

func BenchmarkIsAnagram_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsAnagram_2("anagramanagramanagramanagramanagram", "nagaramnagaramnagaramnagaramnagaramnagaram")
	}
}
