type node struct {
    Val byte
    Next *node
}

func isValid(s string) bool {
    m := map[byte]byte{')':'(', '}':'{', ']':'['}
    var stack *node
    for i := range s{
        if v, ok := m[s[i]]; !ok {
            newNode := node {s[i], stack} // push
            stack = &newNode
        } else {
            if stack == nil || v != stack.Val {
                return false
            }
            stack = stack.Next // pop
        }
    }
    return stack == nil
}
