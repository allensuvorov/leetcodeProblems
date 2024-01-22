// fancy digital root solution
func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	if num % 9 == 0 {
		return 9
	}

	return num % 9
}

// iteration solution
func addDigits(num int) int {
    for num / 10 > 0 {
        sum := 0
        for num > 0 {
            dig := num % 10
            sum += dig
            num /= 10
        }
        num = sum
    }
    return num
}
