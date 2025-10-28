func isSubsequence(s string, t string) bool {
    check := 0 // sequential check list
    for scan := 0; check < len(s) && scan < len(t); scan++ {
        if s[check] == t[scan] {
            check++
        }
    }
    return check == len(s)
}
