func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }
    maxLen := gcd(len(str1), len(str2))
    return str1[:maxLen]
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}


// more intuitive

func gcdOfStrings(str1 string, str2 string) string {
    n, m := len(str1), len(str2)

    if str1 + str2 != str2 + str1 {
        return ""
    }

    for d := min(n, m); d > 0; d-- {
        if n % d == 0 && m % d == 0 {
            return str1[:d]
        }
    }

    return ""
}
