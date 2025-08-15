func mergeAlternately(word1 string, word2 string) string {
    n1, n2 := len(word1), len(word2)
    res := make([]byte, 0, n1 + n2)
    i, j := 0, 0
    for range max(n1, n2) {
        if i < len(word1) {
            res = append(res, word1[i])
            i++
        }
        if j < len(word2) {
            res = append(res, word2[j])
            j++
        }
    }
    return string(res)
}
