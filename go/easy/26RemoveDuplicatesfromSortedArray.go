package main

import (
	"fmt"
)
func removeDuplicates(nums []int) int {
// logic - find unique values and line them up leftside
// this is top solution on leetcode, only one loop, catches only unique, and just roles through the rest
	if len(nums) == 0 {
		return 0
	}
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[length] = nums[i]
			length++
		}
	}
	return length
}
func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}
