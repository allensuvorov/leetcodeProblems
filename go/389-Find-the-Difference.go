func findTheDifference(s string, t string) byte {
    m := make(map[byte]int)
    for i := range s {
        m[s[i]]--
        m[t[i]]++
    }
    m[t[len(t)-1]]++
    for k, v := range m {
        if v != 0 {
            return k
        }
    }
    return t[0]
}
