package Minimum_Size_Subarray_Sum

// 209
// https://leetcode.cn/problems/minimum-size-subarray-sum/
// 双指针 滑动窗口
func MinSubArrayLen(target int, nums []int) int {
	var (
		result = len(nums) + 1
		sum    int
	)

	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target { // 当left==right sum==0 且target>0 因此不用判left<=right
			if result > (right - left + 1) {
				result = (right - left + 1)
			}
			sum -= nums[left]
			left++
		}
	}

	if result == len(nums)+1 {
		return 0
	}
	return result
}
