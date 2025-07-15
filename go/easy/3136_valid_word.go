func isValid(word string) bool {
    // It contains a minimum of 3 characters.
    if len(word) < 3 {
        return false
    }
    includesVowel := false
    includesConsonant := false
    isVowel := map[rune]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
    for _, v := range word {
        // It contains only digits (0-9), and English letters (uppercase and lowercase).
        if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
            return false
        }

        if unicode.IsLetter(v) {
            if isVowel[unicode.ToLower(v)] {
                // It includes at least one vowel.
                includesVowel = true
            } else {
                // It includes at least one consonant.
                includesConsonant = true
            }
        }
    }
    return includesVowel && includesConsonant
}
