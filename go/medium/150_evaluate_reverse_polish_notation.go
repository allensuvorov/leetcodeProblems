func evalRPN(tokens []string) int {
    ops := map[string]func(int, int)int{
        "+": func(a, b int) int{
            return a + b
        },
        "-": func(a, b int) int{
            return a - b
        },
        "*": func(a, b int) int{
            return a * b
        },
        "/": func(a, b int) int{
            if b == 0 {
                return -1
            }
            return a / b
        },
    }

    stack := []int{}
    for _, v := range tokens {
        if _, ok := ops[v]; !ok {
            if num, err := strconv.Atoi(v); err != nil {
                fmt.Println(err)
            } else {
                stack = append(stack, num)
            }
        } else {
            top := len(stack) - 1
            b := stack[top]
            a := stack[top - 1]
            stack = stack[:top]
            stack[len(stack) - 1] = ops[v](a, b)
        }
    }
    return stack[0]
}
