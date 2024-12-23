func isValid(s string) bool {
    stack := []byte{}
    open := map[byte]byte{
        '(':')', 
        '{':'}',
        '[':']',
    }

    for i := range s {
        if _, exists := open[s[i]]; exists { // open paren
            stack = append(stack, s[i])
        } else { // closed paren
            if len(stack) == 0 || open[stack[len(stack)-1]] != s[i] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}
