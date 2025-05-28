func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }

    gcdLen := min(len(str1), len(str2))
    for gcdLen > 1 && (len(str1) % gcdLen != 0 || len(str2) % gcdLen != 0) {
        gcdLen--
    }

    return str1[:gcdLen]
}
