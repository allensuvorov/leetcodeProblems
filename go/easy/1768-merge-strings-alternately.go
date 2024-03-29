func mergeAlternately(word1 string, word2 string) string {
    res := make([]byte, 0, len(word1) + len(word2))
    long := 0
    if len(word1) > len(word2) {
        long = len(word1)
    } else {
        long = len(word2)
    }

    for i := 0; i < long; i ++ {
        if i < len(word1) {
            res = append(res, word1[i])
        }
        if i < len(word2) {
            res = append(res, word2[i])
        }
    }
    
    return string(res)
}
