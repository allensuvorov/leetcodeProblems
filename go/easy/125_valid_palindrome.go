func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l < r {
        if !isAlNu(s[l]) {
            l++
            continue
        }
        if !isAlNu(s[r]) {
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

func isAlNu(b byte) bool {
    if (b >= '0' && b <= '9') || 
        (b >= 'a' && b <= 'z') || 
        (b >= 'A' && b <= 'Z') {
        return true
    }
    return false
}

func toLower(b byte) byte {
    if (b >= 'A' && b <= 'Z') {
        b += 'a' - 'A'
    }
    return b
}
