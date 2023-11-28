func countPrefixes(words []string, s string) int {
    count := 0
    for _, w := range words {
        if len(w) <= len(s) && w == s[:len(w)] {
            count++
        }
    }
    return count
}
