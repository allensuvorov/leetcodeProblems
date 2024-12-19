func isPalindrome(s string) bool {

    isAlphaNumeric := func(b byte) bool {
        r := rune(b)
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            return true
        }
        return false
    }

    toLower := func(b byte) byte {
        r := rune(b)
        r = unicode.ToLower(r)
        return byte(r)
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
