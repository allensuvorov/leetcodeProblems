func countSegments(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
            res++
        }
    }
    return res
}
