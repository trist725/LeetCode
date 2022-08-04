package Reverse_String_II

import "testing"

func TestReverseStr(t *testing.T) {
	println(ReverseStr_1("krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc", 20))
	println(ReverseStr_1("abcdefg", 2))
	println(ReverseStr_1("abcdefg", 1))
}

func BenchmarkReverseStr_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseStr_1("krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc", 20)
	}
}

func BenchmarkReverseStr_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseStr_2("krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc", 20)
	}
}
