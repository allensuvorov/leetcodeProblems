func maxDepth(s string) int {
    max := 0
    stack := make([]byte, 0, len(s))
    for i := range s {
        if s[i] == '(' {
            stack = append(stack, '(')
            if len(stack) > max {
                max = len(stack)
            }
        }
        if s[i] == ')' {
            stack = stack[:len(stack)-1]
        }
    }
    return max
}
