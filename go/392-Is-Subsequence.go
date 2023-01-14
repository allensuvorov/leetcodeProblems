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
