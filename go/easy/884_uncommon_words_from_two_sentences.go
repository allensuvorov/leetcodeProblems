func uncommonFromSentences(s1 string, s2 string) []string {
    m := map[string]int{}

    f := func(s string){
        var b strings.Builder
        for i := range s {
            if s[i] != ' ' {
                b.WriteByte(s[i])
            } else {
                m[b.String()]++
                b.Reset()
            }
        }
        m[b.String()]++
    }
    
    f(s1)
    f(s2)

    ans := []string{}
    for k, v := range m {
        if v == 1 {
            ans = append(ans, k)
        }
    }
    return ans
}
