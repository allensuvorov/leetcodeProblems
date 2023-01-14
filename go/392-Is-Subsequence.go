func isSubsequence(s string, t string) bool {
    j := 0
    hm := map[byte]int{}
    for i := range t {
        hm[t[i]] ++
    }

    for i := 0; i<len(s); i++ {
        if j > len(t) - 1 {
            return false
        }
        
        if hm[s[i]] <= 0 {
            return false
        }

        inSubseq := false
        for j < len(t) {
            if s[i] == t[j] {
                inSubseq = true
                hm[s[i]] --
                j++
                break
            }
            j++
        }
        if !inSubseq {
            return false
        }
    }
    return true
}

// a shorter one
func isSubsequence(s string, t string) bool {
    i,j := 0, 0
    hm := map[byte]int{}
    for i := range t {
        hm[t[i]] ++
    }
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            hm[s[i]] --
            i++
            if i < len(s) && hm[s[i]] <= 0 {
                return false
            }
        }
        j++
    }
    return i == len(s)
}
