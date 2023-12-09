func commonChars(words []string) []string {
    fm1 := make(map[byte]int)
    for i := range words[0] {
        fm1[words[0][i]]++
    }
    for _, w := range words {
        fm2 := make(map[byte]int)
        for i := range w {
            fm2[w[i]]++
        }
        for k, v := range fm1 {
            if fm2[k] < v {
                fm1[k] = fm2[k]
            }
            if fm1[k] == 0 {
                delete(fm1, k)
            }
        }
    }
    res := make([]string, 0, len(fm1))
    for k, v := range fm1 {
        for i := 1; i <= v; i++ {
            res = append(res, string(k))
        }
    }
    return res
}
