func canBeTypedWords(text string, brokenLetters string) int {
    res := 0
    for _, word := range strings.Fields(text) {
        canBeTyped := 1
        for i := range brokenLetters {
            if strings.Contains(word, brokenLetters[i:i+1]) {
                canBeTyped = 0
            }
        }
        
        res += canBeTyped
    }
    return res
}
