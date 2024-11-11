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
    symbolToValue := map[string]int{
        "I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, 
                "XL": 40, "L": 50, "XC": 90, "C": 100,
                "CD": 400, "D": 500, "CM": 900, "M": 1000,
    }

    value := []int{    1,   4,    5,   9,    10,  40,   50,  90,   100, 400,  500, 900,  1000}
    symbol := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
    
    ans := ""
    for i := len(value) - 1; i >= 0; i-- {
        if value[i] <= num {
            ans = symbol[i]
            break
        }
    }

    return ans + intToRoman(num - symbolToValue[ans])
}
