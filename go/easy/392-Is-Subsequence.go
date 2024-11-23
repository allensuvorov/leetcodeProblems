func isSubsequence(s string, t string) bool {
    j := 0
    for i := 0; i < len(t) && j < len(s); i++ {
        if t[i] != s[j] {
            continue
        }
        j++
    }
    return j == len(s)
}
