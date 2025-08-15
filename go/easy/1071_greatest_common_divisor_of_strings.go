func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }

    divLen := min(len(str1), len(str2))
    for divLen >= 1 && !(len(str1) % divLen == 0 && len(str2) % divLen == 0) {
        divLen-- 
    }

    return str1[:divLen]
}
