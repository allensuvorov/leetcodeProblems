func compress(chars []byte) int {
    n := len(chars)
    w := 0
    for l := 0; l < n; {
        // read
        r := l + 1
        for r < n && chars[l] == chars[r] {
            r++
        }
        // write
        chars[w] = chars[l] // write char
        w++
        if (r-l) > 1 { // write char count
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
