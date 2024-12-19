func intToRoman(num int) string {
    symbols := []string {"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
    values := []int     {1,    4,    5,   9,    10,  40,   50,  90,   100, 400, 500, 900, 1000}

    roman := ""

    for i := len(values) - 1; num > 0 && i >= 0; {
        if values[i] <= num {
            roman += symbols[i]
            num -= values[i]
        } else {
            i--
        }
    }

    return roman
}

func intToRoman(num int) string {
    if num == 0 {
        return ""
    }

    values := []int     {1,    4,    5,   9,    10,  40,   50,  90,   100, 400, 500, 900, 1000}
    symbols := []string {"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
    
    i := len(values) - 1
    for i >= 0 {
        if values[i] <= num {
            break
        }
        i--
    }

    return symbols[i] + intToRoman(num - values[i])
}
