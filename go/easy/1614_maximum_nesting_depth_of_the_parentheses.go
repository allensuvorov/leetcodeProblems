func maxDepth(s string) int {
    max := 0
    count := 0
    var stack *StackNode
    for i := range s {
        if s[i] == '(' {
            newNode := StackNode{'(', stack}
            stack = &newNode
            count++
            if count > max {
                max = count
            }
        }
        if s[i] == ')' {
            stack = stack.Next
            count--
        }
    }
    return max
}

type StackNode struct{
    Val byte
    Next *StackNode
}
