func reverseParentheses(s string) string {
    temp := []byte{}
    stack := []byte{}
    for i := range s {
        if s[i] == ')' {
            fmt.Println(len(stack))
            for stack[len(stack) - 1] != '(' {
                top := len(stack) - 1
                temp = append(temp, stack[top])
                stack = stack[:top]
            }
            stack = stack[:len(stack) - 1]
            stack = append(stack, temp...)
            temp = nil
        } else {
            stack = append(stack, s[i]) 
        }
    }
    return string(stack)
}
