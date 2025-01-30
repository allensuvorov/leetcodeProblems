func compress(chars []byte) int {
    plant := 0
    for l := 0; l < len(chars); {
        // read a group
        r := l + 1
        for r < len(chars) && chars[r] == chars[l] {
            r++
        }
        // write compressed group info
        chars[plant] = chars[l]
        plant++

        if r - l > 1 {
            s := strconv.Itoa(r - l)
            for i := range s {
                chars[plant] = s[i]
                plant++
            }
        }
        l = r
    }
    return plant
}
