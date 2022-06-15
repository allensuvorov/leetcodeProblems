package main

import (
	"fmt"
)

func singleNumber (nums []int) int {
  result := nums[0]
  for i := 1; i<len(nums);i++ {
    result ^= nums[i]
  }
  return result
}

func main() {
  nums := []int {1,1,2,2,3}
  fmt.Println(singleNumber(nums))
	fmt.Println("Hello, World!")
}
