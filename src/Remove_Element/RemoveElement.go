package Remove_Element

func RemoveElement_1(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

//快慢指针
func RemoveElement_2(nums []int, val int) int {
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow]
	return len(nums)
}
