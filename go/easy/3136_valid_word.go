func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
    
    hasVowel := false
    hasConsonant := false
    isVowel := map[rune]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}

    for _, v := range word {
        if unicode.IsLetter(v) {
            if isVowel[unicode.ToLower(v)] {
                hasVowel = true
            } else {
                hasConsonant = true
            }
        } else if !unicode.IsDigit(v) {
            return false
        }
    }
    return hasVowel && hasConsonant
}
