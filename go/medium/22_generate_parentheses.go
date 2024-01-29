func generateParenthesis(n int) []string {
    res := []string{}
    stack := []byte{}
    var dfs func(op, cl int)
    dfs = func(op, cl int) {
        if cl == n  {
            res = append(res, string(stack))
            return 
        }
        if op < n {
            stack = append(stack, '(')
            dfs(op+1, cl)
            stack = stack[:len(stack) - 1]
        }
        if op > cl {
            stack = append(stack, ')')
            dfs(op, cl+1)
            stack = stack[:len(stack) - 1]
        }
    }
    dfs(0, 0)
    return res
}
