func lexicalOrder(n int) []int {
    ans := make([]int, 0, n)
    var dfs func(int)
    dfs = func(cur int) {
        if cur > n {
            return
        }
        ans = append(ans, cur)
        for i := 0; i < 10; i++ {
            dfs(cur * 10 + i)
        }
    }
    for v := 1; v <= 9; v ++ { 
        dfs(v)
    }
    return ans
}
