package Squares_Of_A_Sorted_Array

func sortedSquares(nums []int) []int {
	var (
		rIdx   = len(nums) - 1
		result = make([]int, len(nums))

		left  int
		right = len(nums) - 1
	)

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[rIdx] = nums[left] * nums[left]
			left++
		} else {
			result[rIdx] = nums[right] * nums[right]
			right--
		}
		rIdx--
	}

	return result
}
