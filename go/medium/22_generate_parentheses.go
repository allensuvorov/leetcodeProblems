func generateParenthesis(n int) []string {
    res := []string{}
    var dfs func(op, cl, rem int, comb []byte)
    dfs = func(op, cl, rem int, comb []byte) {
        if rem == 0 {
            res = append(res, string(comb))
            return 
        }
        if op < n {
            dfs(op+1, cl, rem-1, append(comb, '('))
        }
        if op > cl {
            dfs(op, cl+1, rem-1, append(comb, ')'))
        }
    }
    dfs(0, 0, n*2, []byte{})
    return res
}
