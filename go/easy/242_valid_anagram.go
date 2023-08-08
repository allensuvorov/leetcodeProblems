package easy

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    byteCount := map[byte]int{}
    for i := range s {
        byteCount[s[i]]++
        byteCount[t[i]]--
    }

    for i := range byteCount {
        if byteCount[i] != 0 {
            return false
        }
    }
    return true
}
