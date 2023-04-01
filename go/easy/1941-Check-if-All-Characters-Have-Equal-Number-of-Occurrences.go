func areOccurrencesEqual(s string) bool {
    hm := map[byte]int{}
    for i := range s {
        hm[s[i]]++
    }   
    standard := hm[s[0]]
    for i := range hm {
        if hm[i] != standard {
            return false
        }
    }
    return true
}
