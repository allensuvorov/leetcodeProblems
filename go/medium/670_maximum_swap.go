func maximumSwap(num int) int {
	digits := intByDigits(num)
    last := make([]int, 10)
    
	for i := range digits {
		last[digits[i]] = i
	}

	for i, digit := range digits {
		for j := len(last) - 1; j > digit; j-- {
			if last[j] > i {
				digits[i], digits[last[j]] = digits[last[j]], digits[i]
				return digitsToInt(digits)
			}
		}
	}
	return num
}

func intByDigits(num int) []int {
	digits := []int{}
	for x := num; x > 0; x /= 10 {
		digit := x % 10
		digits = append(digits, digit)
	}

	// reverse digits
	for l, r := 0, len(digits)-1; l < r; {
		digits[l], digits[r] = digits[r], digits[l]
		l++
		r--
	}
	return digits
}

func digitsToInt(digits []int) int {
	ans := 0
	for _, v := range digits {
		ans *= 10
		ans += v
	}
	return ans
}
