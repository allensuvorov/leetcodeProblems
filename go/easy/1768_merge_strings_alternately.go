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

// channels practice, receive from two channels
func mergeAlternately(word1 string, word2 string) string {
    c1 := make(chan byte)
    c2 := make(chan byte)
    go letters(word1, c1)
    go letters(word2, c2)
    res := make([]byte, 0, len(word1) + len(word2))
    for range cap(res) {
        if v, ok := <- c1; ok {
            res = append(res, v)
        }
        if v, ok := <- c2; ok {
            res = append(res, v)
        }
    }
    return string(res)
}

func letters(word string, c chan byte) {
    for i := range word {
        c <- word[i]
    }
    close(c)
}
