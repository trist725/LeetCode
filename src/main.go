package main

import (
	"fmt"
	t "LeetCode/src/two_sum"
)

func main(){
	nums := []int{3,2,3}
	fmt.Println(nums[0:])
	target := 6
	output := t.TwoSum(nums, target)
	fmt.Println(output)
}