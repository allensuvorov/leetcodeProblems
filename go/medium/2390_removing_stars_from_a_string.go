func removeStars(s string) string {
    var stack []byte
    for i := range s {
        if s[i] != '*' {
            stack = append(stack, s[i]) // push
        } else {
            stack = stack[:len(stack)-1] // pop
        }
    }
    return string(stack)
}
