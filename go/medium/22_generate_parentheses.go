func generateParenthesis(n int) []string {
    res := []string{}
    var dfs func(op, cl int, stack []byte)
    dfs = func(op, cl int, stack []byte) {
        if cl == n  {
            res = append(res, string(stack))
            return 
        }
        if op < n {
            dfs(op+1, cl, append(stack, '('))
        }
        if op > cl {
            dfs(op, cl+1, append(stack, ')'))
        }
    }
    dfs(0, 0, []byte{})
    return res
}
