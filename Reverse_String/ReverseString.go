package Reverse_String

// https://leetcode.cn/problems/reverse-string/

func reverseString(s []byte) {
	var (
		left  = 0
		right = len(s) - 1
	)

	for left != right {
		// 1 <= s.length <= 105
		s[left], s[right] = s[right], s[left]
		left++
		if left == right {
			break
		}
		right--
	}
}
