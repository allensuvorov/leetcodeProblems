func isIsomorphic(s string, t string) bool {
    mapping_s_t := make(map[byte]byte)
    mapping_t_s := make(map[byte]byte)

    for i := range s {
        _, ok1 := mapping_s_t[s[i]]
        _, ok2 := mapping_t_s[t[i]]
        if !ok1 && !ok2 {
            mapping_s_t[s[i]] = t[i]
            mapping_t_s[t[i]] = s[i]
        } else if mapping_s_t[s[i]] != t[i] || mapping_t_s[t[i]] != s[i] {
                return false
        }
    }
    return true
}
