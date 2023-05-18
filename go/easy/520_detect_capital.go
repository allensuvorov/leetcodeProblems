func detectCapitalUse(word string) bool {
    if len(word) == 1 {
        return true
    }
    var checkFunc func(r rune) bool
    if unicode.IsUpper(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
        checkFunc = unicode.IsUpper
    } else {
        checkFunc = unicode.IsLower
    }

    for i := 1; i < len(word); i++ {
        if !checkFunc(rune(word[i])) {
            return false
        }
    }
    return true
}
