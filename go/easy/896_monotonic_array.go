package main

var a1 = []int{1, 2, 3} // true
var a2 = []int{1, 2, 1} // false

func isMonotonic(nums []int) bool {
	// first, find diff and save sign (+/-)
	// second, find diff with opposite sign (+/-)
	sign := ""
	for i := 0; i < len(nums)-1; i++ {
		if sign == "" {
			if diff := nums[i+1] - nums[i]; diff != 0 {
				sign = getSign(diff)
			}
		} else {
			if diff := nums[i+1] - nums[i]; diff != 0 {
				if sign != gitSign(diff) {
					return false
				}
			}
		}
	}
	return true
}

func getSign(n int) string {
	sign := ""
	if n < 0 {
		sign = "-"
	} else {
		sign = "+"
	}
	return sign
}

// 4 cases
// 1 all numbers are the same
// 2 one or more increase
// 3 one or more decrease
// 4 both inc and dec exist - chech if that is the case or not
