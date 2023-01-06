func findTheDifference(s string, t string) byte {
    m := make(map[byte]int)
    for i := range s {
        m[t[i]]++
        m[s[i]]--
        if m[s[i]] == 0 {
            delete(m, s[i])
        }
    }
    m[t[len(t)-1]]++
    var result byte
    for k, v := range m {
        if v != 0 {
            result = k
        }
    }
    return result
}
