func mergeAlternately(word1 string, word2 string) string {
    res := make([]byte, 0, len(word1) + len(word2))
    for i := range max(len(word1), len(word2)) {
        if i < len(word1) {
            res = append(res, word1[i])
        }
        if i < len(word2) {
            res = append(res, word2[i])
        }
    }
    return string(res)
}

// Big O (n^2) In Go strings are immutable. Every concat is a O(n).
func mergeAlternately(word1 string, word2 string) string {
    res := ""
    
    for i := range max(len(word1), len(word2)) {
        
        if i < len(word1) {
            res += string(word1[i])
        }
        
        if i < len(word2) {
            res += string(word2[i])
        }
    }
    return res
        
}
