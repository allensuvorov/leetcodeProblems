func numTrees(n int) int {
    dp := []int{0:1, 1:1, 20: 0}

    for i := 2; i < n + 1; i++ {
        sum := 0
        for r := range i {
            l := i - (r + 1)
            sum += dp[l] * dp[r]
        }
        dp[i] = sum
    }

    return dp[n]
}
