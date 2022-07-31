package Intersection_of_Two_Arrays

// https://leetcode.cn/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	var (
		count  = make(map[int]struct{})
		result []int
	)

	for _, v := range nums1 {
		count[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := count[v]; ok {
			result = append(result, v)
			// 后面重复的不会再进结果数组
			delete(count, v)
		}
	}

	return result
}
