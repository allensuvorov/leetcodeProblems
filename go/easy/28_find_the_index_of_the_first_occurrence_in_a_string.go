func strStr(haystack string, needle string) int {
    for i := 0; i < len(haystack) - len(needle) + 1; i++ {
        j := 0
        for j < len(needle) && haystack[i+j] == needle[j] {
            j++
        }
        if j == len(needle) {
            return i
        }        
    }
    return -1
}
