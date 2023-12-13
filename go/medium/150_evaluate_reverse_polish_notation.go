func evalRPN(tokens []string) int {
    var stack *node
    for _, v := range tokens {
        if v == "+" || v == "-" || v == "*" || v == "/" {
            stack.Next.Val = ops[v](stack.Next.Val, stack.Val)
            stack = stack.Next
        } else {
            num, err := strconv.Atoi(v)
            if err != nil {
		        fmt.Println(err)
	        }  
            newNode := node{num, stack}
            stack = &newNode
        }
    }
    return stack.Val
}

type node struct {
    Val int
    Next *node
}

var ops = map[string]func(int, int) int {
        "+": func(a, b int) int { return a + b },
        "-": func(a, b int) int { return a - b },
        "*": func(a, b int) int { return a * b },
        "/": func(a, b int) int { return a / b },
    }
