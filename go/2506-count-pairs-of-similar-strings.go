func similarPairs(words []string) int {
    res := 0
    collections := map[string]int{}
    for _, w := range words {
        abc := make([]byte, 26)
        for i := range w {
            abc [w[i]- 'a'] = w[i]
        }

        abcString := string(abc)
        if _, ok := collections[abcString]; ok {
            res += collections[abcString]
        }
        collections[abcString]++
    }
    return res
}
