func makeFancyString(s string) string {
    res := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        if i == 0 || s[i] != s[i-1] || i == len(s) - 1 || s[i] != s[i+1] {
            res = append(res, s[i])
        }
    }
    return string(res)
}
