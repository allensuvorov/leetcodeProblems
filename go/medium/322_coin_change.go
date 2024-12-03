func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)

    for i := 1; i < len(dp); i++ {
        best := math.MaxInt
        for _, j := range coins {
            if j == i {
                best = 1
                break
            }
            if j < i && dp[j] != -1 && dp[i - j] != -1 {
                cur := dp[j] + dp[i - j] // cur best
                best = min(best, cur)
            }
        }

        if best == math.MaxInt {
            dp[i] = -1
        } else {
            dp[i] = best
        }
    }
    fmt.Println(dp)
    return dp[len(dp)-1]
}
