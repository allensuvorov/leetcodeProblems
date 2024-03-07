func maximum69Number (num int) int {
    stack := []int{}
    res := 0
    for tmp := num; tmp > 0; tmp /= 10 {
        stack = append(stack, tmp % 10)
    }
    
    for i := len(stack) - 1; i >= 0; i-- {
        if stack[i] == 6 {
            stack[i] = 9
            break
        }
    }

    for len(stack) > 0 {
        top := len(stack) - 1
        res = res*10 + stack[top]
        stack = stack[:top]   
    }

    return res
}
