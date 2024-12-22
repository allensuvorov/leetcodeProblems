func isIsomorphic(s string, t string) bool {
    return isMapping(s,t) && isMapping(t,s)
}

func isMapping(s, t string) bool {
    st := make(map[byte]byte)
    for i := range s {
        if v, ok := st[s[i]]; ok && v != t[i]{
            return false
        } else {
            st[s[i]] = t[i]
        }
    }
    return true
}
