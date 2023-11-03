func isPrefixString(s string, words []string) bool {
    sl := len(s)
    i := 0
    for _, k := range words {
        wl := len(k)
        if i + wl > sl || k != s[i:i+wl] {
            return false
        }
        i += wl
        if i == sl {
            return true
        }
    }
    return false
}
