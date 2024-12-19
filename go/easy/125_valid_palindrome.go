func isPalindrome(s string) bool {

    isAlphaNumeric := func(b byte) bool {
        r := rune(b)
        return unicode.IsLetter(r) || unicode.IsDigit(r)
    }

    toLower := func(b byte) rune {
        r := rune(b)
        return unicode.ToLower(r)
    }

    for l, r := 0, len(s) - 1; l < r; {
        if !isAlphaNumeric(s[l]) {
            l++
            continue
        }
        if !isAlphaNumeric(s[r]) {
            r--
            continue
        }

        if toLower(s[l]) != toLower(s[r]) {
            return false
        }
        l++
        r--
    }
    
    return true
}
