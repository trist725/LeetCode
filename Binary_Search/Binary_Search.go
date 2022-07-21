package Binary_Search

// 704
// https://leetcode.cn/problems/binary-search/

// 暴力枚举
func BinarySearch_1(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

// 二分查找
func BinarySearch_2(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	return -1
}
