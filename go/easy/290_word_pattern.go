func wordPattern(pattern string, s string) bool {
    bm := map[byte]string{}
    wm := map[string]byte{}
    words := strings.Fields(s)
    if len(pattern) != len(words) {
        return false
    }
    for i := 0; i < len(pattern); i++ {
        if w, ok := bm[pattern[i]]; ok {
            if w != words[i] {
                return false
            }
        } else if b, ok := wm[words[i]]; ok{
            if b != pattern[i] {
                return false
            }
        } else {
            bm[pattern[i]] = words[i]
            wm[words[i]] = pattern[i]
        }
    }
    return true
}
