func isIsomorphic(s string, t string) bool {
    ab := map[byte]byte{s[0]:t[0]}
    ba := map[byte]byte{t[0]:s[0]}

    for i := range s {
        if ab[s[i]] != t[i] || ba[t[i]] != s[i] {
            return false
        }
        ab[s[i]] = t[i]
        ba[t[i]] = s[i]
    }

    return true
}
