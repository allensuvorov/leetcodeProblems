func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }
    maxLen := 0
    for i := range min(len(str1), len(str2)) {
        if len(str1) % (i+1) == 0 && len(str2) % (i+1) == 0 {
            maxLen = i + 1
        }
    }
    return str1[:maxLen]
}
