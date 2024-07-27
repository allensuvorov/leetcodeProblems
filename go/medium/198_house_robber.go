func rob(nums []int) int {
    dp := make([]int, len(nums))
    maxMoney := 0
    for i, stash := range nums {
        
        a := 0
        b := 0
        if i >= 2 {
            a = dp[i-2]
        }
        if i >= 3 {
            b = dp[i-3]
        }
        bestCur := max(a, b) + stash
        maxMoney = max(maxMoney, bestCur)
        dp[i] = bestCur
    }
    return maxMoney
}
