func truncateSentence(s string, k int) string {
    var b strings.Builder
    count := 0
    for i := range s{
        if s[i] == ' ' {
            count++
        }
        if count == k {
            break
        }
        b.WriteByte(s[i])
    }
    return b.String()
}
