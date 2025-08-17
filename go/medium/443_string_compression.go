func compress(chars []byte) int {
    r := 0 // read index
    w := 0 // write index
    for r < len(chars) {
        // write the current group character
        chars[w] = chars[r]
        
        // read the group length
        groupLen := 0
        for r < len(chars) && chars[r] == chars[w] {
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
