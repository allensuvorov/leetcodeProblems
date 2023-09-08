func isValid(s string) bool {
    m := map[byte]byte{')':'(', '}':'{', ']':'['}
    stack := []byte{}
    for i := range s{
        if v, ok := m[s[i]]; !ok {
            stack = append(stack, s[i])
        } else {
            if len(stack) == 0 || v != stack[len(stack)-1] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
