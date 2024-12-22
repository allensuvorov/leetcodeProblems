func isIsomorphic(s string, t string) bool {
    st := make(map[byte]byte)
    ts := make(map[byte]byte)
    for i := range s {
        if v, ok := st[s[i]]; ok && v != t[i]{
            return false
        } else {
            st[s[i]] = t[i]
        }
        if v, ok := ts[t[i]]; ok && v != s[i] {
            return false
        } else {
            ts[t[i]] = s[i]
        }
    }
    return true
}

