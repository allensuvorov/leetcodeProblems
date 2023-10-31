func minLength(s string) int {
    size := len(s)
    pairs := map[byte]byte{'A':'B', 'C':'D'}
    count := 0
    for i := 0; i < len(s) - 1; {
        if pairs[s[i]] == s[i+1] {
            count += 2
            first := s[:i]
            second := ""
            if i + 1 < len(s) {
                second = s[i+2:]
            }
            s = first + second
            fmt.Println(s,i)
            if i > 0 {
                i--
            }
        } else {
            i++
        }
    }
    return size - count
}
