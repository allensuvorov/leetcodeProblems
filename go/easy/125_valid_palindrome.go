func isPalindrome(s string) bool {
    alphaNumeric := func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            return unicode.ToLower(r)
        }
        return -1
    }
    s = strings.Map(alphaNumeric, s)
    for l, r := 0, len(s) - 1; l < r; {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}
