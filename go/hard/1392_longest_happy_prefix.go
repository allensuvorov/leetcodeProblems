func longestPrefix(s string) string {
    ans := ""
    for i := 1; i < len(s); i++ {
        if strings.HasSuffix(s, s[:i]) {
            ans = s[:i]
        }
    }
    return ans
}
