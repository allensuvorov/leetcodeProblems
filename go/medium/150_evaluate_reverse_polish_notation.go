func evalRPN(tokens []string) int {
    stack := []int{}
    for _, v := range tokens {
        top := len(stack) - 1
        if _, ok := ops[v]; ok {
            stack[top-1] = ops[v](stack[top-1], stack[top])
            stack = stack[:top]
        } else {
            num, err := strconv.Atoi(v)
            if err != nil {
                fmt.Println(err)
            }
            stack = append(stack, num)
        }
    }
    return stack[0]
}

var ops = map[string]func(int, int) int {
    "*": func(a, b int) int { return a * b },
    "/": func(a, b int) int { return a / b },
    "+": func(a, b int) int { return a + b },
    "-": func(a, b int) int { return a - b },
}
