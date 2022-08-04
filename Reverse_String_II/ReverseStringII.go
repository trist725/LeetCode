package Reverse_String_II

import "LeetCode/Reverse_String"

// https://leetcode.cn/problems/reverse-string-ii/

func ReverseStr_1(s string, k int) string {
	if k == 1 {
		return s
	}

	var (
		times = len(s) / (2 * k)
		ss    = []byte(s) // Reverse_String.ReverseString可以直接修改[]byte
	)

	for i := 0; i < times; i++ {
		Reverse_String.ReverseString(ss[i*2*k : i*2*k+k])
	}
	// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样
	if (len(s) - times*2*k) >= k {
		Reverse_String.ReverseString(ss[times*2*k : times*2*k+k])
	} else {
		// 如果剩余字符少于 k 个，则将剩余字符全部反转
		Reverse_String.ReverseString(ss[times*2*k:])
	}
	return string(ss)
}

// 官方题解
func ReverseStr_2(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := t[i:min(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
