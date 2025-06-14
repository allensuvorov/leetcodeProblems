func compress(chars []byte) int {
    w := 0 // writer index
    r := 0 // reader index
    for r < len(chars) {
        chars[w] = chars[r] // new group
        groupLen := 0
        for r < len(chars) && chars[w] == chars[r] {
            groupLen++
            r++
        }
        w++

        if groupLen > 1 {
            s := strconv.Itoa(groupLen)
            for i := range s {
                chars[w] = s[i]
                w++
            }
        }
    }
    return w
}
