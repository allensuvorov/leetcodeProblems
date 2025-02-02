func compress(chars []byte) int {
    w := 0
    for l := 0; l < len(chars); {
        // read, measure 
        r := l + 1
        for r < len(chars) && chars[r] == chars[l] {
            r++
        }
        // write
        chars[w] = chars[l]
        w++
        if (r - l) > 1 {
            s := strconv.Itoa(r - l)
            for i := range s {
                chars[w] = s[i]
                w++
            }
        }
        l = r
    }
    return w
}
