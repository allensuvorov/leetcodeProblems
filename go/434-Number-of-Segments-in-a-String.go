func countSegments(s string) int {
    res := 0
    const space = byte(' ')
    prev := space
    for i := 0; i < len(s); i++ {
        if s[i] != space && prev == space {
            res++
        }
        prev = s[i]
    }
    return res
}
