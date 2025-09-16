func canBeTypedWords(text string, brokenLetters string) int {
    brokenLettersSet := make(map[rune]bool, len(brokenLetters))
    for _, v := range brokenLetters {
        brokenLettersSet[v] = true
    }

    res := 0
    canBeTyped := 1
    for i, char := range text {
        if char != ' ' {
            if brokenLettersSet[char] {
                canBeTyped = 0
            }
        } 
        if char == ' ' || i == len(text) - 1 {
            res += canBeTyped
            canBeTyped = 1
        }
    }
    return res
}
