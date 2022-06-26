package main

import (
	"fmt"
)

func isPossible(target []int) bool {
	max := 0
	sum := 0
	maxI := 0
	rest := 0 // rest of array, excluding max
	if len(target) == 1 {
		if target[0] == 1 {
			return true
		} else {
			return false
		}
	}

	for _, num := range target {
		sum += num
	}

	for {
		max = 0

		for i, num := range target {
			if num > max {
				max = num
				maxI = i
			}
		}
		fmt.Println(target, sum, max, maxI)
		rest = sum - max

		if max == 1 || rest == 1 {
			return true
		}

		if max < rest+1 || max%rest == 0 {
			return false
		} else {
			target[maxI] = max % rest // replace max with remainder
		}
		sum -= max - target[maxI]
	}
}

func main() {

	fmt.Println(isPossible([]int{1, 1}))
	fmt.Println(isPossible([]int{2, 900000001}))
	fmt.Println(isPossible([]int{2}))
}
