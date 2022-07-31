package Find_Common_Characters

// https://leetcode.cn/problems/find-common-characters/

func commonChars(words []string) []string {
	var (
		count  = make(map[rune]int)
		result []string
	)

	// len(words)>=1
	// 初始化第一个string
	for _, ch := range words[0] {
		count[ch]++
	}
	for _, w := range words[1:] {
		countIn := make(map[rune]int)
		for _, ch := range w {
			countIn[ch]++
		}
		for ch, c := range count {
			if countIn[ch] < c {
				count[ch] = countIn[ch]
			}
		}
	}

	for ch, c := range count {
		for i := 0; i < c; i++ {
			result = append(result, string(ch))
		}
	}

	return result
}
