func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
  
    invalidSymbols := map[rune]bool {
        '@': true, 
        '#': true,
        '$': true,
    }
  
    vowels := map[rune]bool {
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }

    var includesVowel, includesConsonant bool

    for _, v := range word {
        v = unicode.ToLower(v)
        if invalidSymbols[v] {
            return false
        }

        if vowels[v] {
            includesVowel = true
        }

        if unicode.IsLetter(v) && !vowels[v] {
            includesConsonant = true
        }
    }

    return includesVowel && includesConsonant
}
