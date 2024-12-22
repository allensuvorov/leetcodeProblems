func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)
    if len(words) != len(pattern) {
        return false
    }
    lw := make(map[byte]string)
    wl := make(map[string]byte)

    for i := range words {
        if v, ok := lw[pattern[i]]; ok && v != words[i] {
            return false
        } else {
            lw[pattern[i]] = words[i]
        }
        if v, ok := wl[words[i]]; ok && v != pattern[i] {
            return false
        } else {
            wl[words[i]] = pattern[i]
        }
    }
    return true
}
