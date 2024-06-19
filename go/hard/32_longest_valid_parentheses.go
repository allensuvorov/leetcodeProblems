func longestValidParentheses(s string) int {
    stack := []int{-1}
    maxLen := 0
    for i, c := range s {
        if c == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack) - 1] // pop
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                maxLen = max(maxLen, i - stack[len(stack) - 1])
            }
        }
    }
    return maxLen
}
