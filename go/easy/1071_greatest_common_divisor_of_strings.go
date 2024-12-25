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
