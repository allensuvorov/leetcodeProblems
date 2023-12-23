func removeStars(s string) string {
    stack := make([]byte, 0, len(s))
    for i := range s {
        if s[i] == '*' {
            stack = stack[:len(stack) - 1] // pop
        } else {
            stack = append(stack, s[i]) // push
        }
    }
    return string(stack)
}
