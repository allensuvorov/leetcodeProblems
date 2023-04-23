func isValid(s string) bool {
    stack := []byte{}
    opn := map[byte]bool{
        '(': true,
        '{': true,
        '[': true,
    }
    cls := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    for i := range s {
        if opn[s[i]] {
            stack = append(stack, (s[i]))
        } else {
            if len(stack) == 0 {
                return false
            }
            if stack[(len(stack)-1):][0] != cls[s[i]] {
                return false
            }
            stack = stack[:(len(stack)-1)]
        }
    }
    if len(stack) != 0 {
        return false
    }
    return true
}
