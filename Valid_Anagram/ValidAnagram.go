package Valid_Anagram

// https://leetcode.cn/problems/valid-anagram/

func IsAnagram_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var (
		countS = make(map[rune]int32)
		countT = make(map[rune]int32)
	)

	for _, v := range s {
		countS[v]++
	}
	for _, v := range t {
		countT[v]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}

func IsAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var (
		count = make(map[rune]int32, len(s))
	)

	for _, v := range s {
		count[v]++
	}
	for _, v := range t {
		count[v]--
		// 因为len相等 减完最后一定全是0 否则不符合
		if count[v] < 0 {
			return false
		}
	}

	return true
}
