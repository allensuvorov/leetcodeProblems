func isIsomorphic(s string, t string) bool {
    return isMatching(s,t) && isMatching(t,s)
}

func isMatching(s string, t string) bool {
    ab := map[byte]byte{}
    for i := range s {
        if v, ok := ab[s[i]]; !ok {
            ab[s[i]] = t[i]
        } else {
            if v != t[i] {
                return false
            }
        }
    }
    return true
}
