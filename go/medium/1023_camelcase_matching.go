func camelMatch(queries []string, pattern string) []bool {
    res := []bool{}
    for _, word := range queries {
        i := 0
        for _, v := range word {
            if i < len(pattern) && v == rune(pattern[i]) {
                i++
            } else if unicode.IsUpper(v) {
                i = -1
                break
            }
        }
        res = append(res, i == len(pattern))
    }
    return res
}
