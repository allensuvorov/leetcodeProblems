func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)

    for a := 1; a < len(dp); a++ {
        dp[a] = math.MaxInt32
        for _, c := range coins {
            if c <= a {
                cur := 1 + dp[a - c] // cur best
                dp[a] = min(dp[a], cur)
            }
        }
    }
    if dp[amount] == math.MaxInt32 {
        return -1
    }
    return dp[amount]
}
