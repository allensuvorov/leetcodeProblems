package easy

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := map[byte]int{}
    for i := range s {
        m[s[i]]++
        m[t[i]]--
        if m[s[i]] == 0 {
            delete(m, s[i])
        }
        if m[t[i]] == 0 {
            delete(m, t[i])
        }
    }
    
    return len(m) == 0
}
