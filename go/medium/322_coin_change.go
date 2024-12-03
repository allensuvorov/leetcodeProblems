func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)

    for a := 1; a < len(dp); a++ {
        dp[a] = amount + 1
        for _, c := range coins {
            if a - c >= 0 {
                cur := 1 + dp[a - c] // cur best
                dp[a] = min(dp[a], cur)
            }
        }
    }
    if dp[amount] == amount + 1 {
        return -1
    }
    return dp[amount]
}
