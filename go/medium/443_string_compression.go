func compress(chars []byte) int {
    n := len(chars)
    w := 0
    for l := 0; l < n; {
        r := l + 1
        for r < n && chars[r] == chars[l] {
            r++
        }
        chars[w] = chars[l]
        w++
        if (r-l) > 1 {
            s := strconv.Itoa(r-l)
            for i := range s {
                chars[w] = s[i]
                w++
            }
        }
        l = r
    }
    return w
}
