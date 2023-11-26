func makeFancyString(s string) string {
    limit := 2
    res := []byte{}
    l := 0
    for r := range s {
        if s[r] != s[l] {
            l = r 
        }
        if r - l < limit {
            res = append(res, s[r])
        }
    }
    return string(res)
}
