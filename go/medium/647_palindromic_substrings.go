func countSubstrings(s string) int {
    count := 0
    for i := range s {
    // check odd lenths
        for l, r := i, i; l >= 0 && r < len(s) && s[l]==s[r]; {
            count++
            l--
            r++
        }
    
    // check all even lenths
        for l, r := i, i+1; l >= 0 && r < len(s) && s[l]==s[r]; {
            count++
            l--
            r++       
        }
    }
    return count
}
