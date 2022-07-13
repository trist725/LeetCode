package main

import (
	"LeetCode/src/two_sum"
	"fmt"
)

func main() {
	//two_sum
	{
		nums := []int{3, 2, 3}
		fmt.Println(nums[0:])
		target := 6

		output := two_sum.TwoSum_1(nums, target)
		fmt.Println(output)
		output = two_sum.TwoSum_2(nums, target)
		fmt.Println(output)
	}

}
