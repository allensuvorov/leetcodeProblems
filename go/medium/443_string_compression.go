func compress(chars []byte) int {
    w := 0
    for r := 0; r < len(chars); {
        // read lenth group
        l := r
        for r < len(chars) && chars[l] == chars[r] { // end of group, r is out of group
            r++
        }

        // write group
        chars[w] = chars[l]
        groupLength := r - l
        w++
        if groupLength > 1 {
            s := strconv.Itoa(groupLength)
            for i := range s {
                chars[w] = s[i]
                w++
            }
        }
    }
    return w
}
