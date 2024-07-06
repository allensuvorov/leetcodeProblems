func numSquares(n int) int {
    ans := n
    var dfs func(rem, cur int)
    dfs = func(rem, cur int) {
        if rem == 0 {
            ans = min(ans, cur)
            return
        }
        for i := 1; i*i <= rem; i++ {
            sq := i*i
            dfs(rem % sq, cur + rem / sq)
        }
    }
    dfs(n, 0)
    return ans
}
