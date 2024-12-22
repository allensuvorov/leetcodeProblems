func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    counts := make([]int, 26)
    for i := range s {
        counts[s[i] - 'a']++
        counts[t[i] - 'a']--
    }

    for _, v := range counts {
        if v != 0 {
            return false
        }
    }

    return true
}
