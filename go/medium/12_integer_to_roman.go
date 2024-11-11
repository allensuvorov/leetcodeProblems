func intToRoman(num int) string {
	value := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	symbol := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	ans := ""
	for num > 0 {
		i := 0
		for i = len(value) - 1; i >= 0; i-- {
			if value[i] <= num {
				break
			}
		}
		num -= value[i]
		ans += symbol[i]
	}

	return ans
}

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}

	value := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	symbol := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	i := 0
	for i = len(value) - 1; i >= 0; i-- {
		if value[i] <= num {
			break
		}
	}

	return symbol[i] + intToRoman(num-value[i])
}
