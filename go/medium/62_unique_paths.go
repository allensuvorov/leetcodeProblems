func uniquePaths(m int, n int) int {
    dp := make([]int, n)

    dp[0] = 1

    for range m {
        for c := 1; c < n; c++ {
            dp[c] += dp[c-1]
        }
    }

    return dp[len(dp) - 1]
}
