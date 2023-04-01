func countGoodSubstrings(s string) int {
    result := 0
    for i := 0; i+2 < len(s); i ++ {
        if (s[i] != s[i+1] && s[i+1] != s[i+2] && s[i] != s[i+2]) {
            result ++
        }
    }
    return result
}
