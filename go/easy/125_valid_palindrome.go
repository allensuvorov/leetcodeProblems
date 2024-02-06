func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    l, r := 0, len(s) - 1
    for l < r {
        if !unicode.IsNumber(rune(s[l])) && !unicode.IsLetter(rune(s[l])) {
            l++
            continue
        }
        if !unicode.IsNumber(rune(s[r])) && !unicode.IsLetter(rune(s[r])) {
            r--
            continue
        }
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}
