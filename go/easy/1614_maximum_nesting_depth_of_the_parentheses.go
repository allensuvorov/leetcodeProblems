func maxDepth(s string) int {
    max := 0
    depth := 0
    for i := range s {
        if s[i] == '(' {
            depth++
            if depth > max {
                max = depth
            }
        }
        if s[i] == ')' {
            depth--
        }
    }
    return max
}
