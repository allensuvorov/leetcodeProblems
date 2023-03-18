func truncateSentence(s string, k int) string {
    res := []byte{}
    count := 0
    for i := range s{
        if s[i] == ' ' {
            count++
        }
        if count == k {
            break
        }
        res = append(res, s[i])
    }
    return string(res)
}
