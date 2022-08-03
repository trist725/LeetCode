package Ransom_Note

// https://leetcode.cn/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int, len(magazine))

	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		if m[v] > 0 {
			m[v]--
		} else {
			return false
		}
	}
	return true
}
