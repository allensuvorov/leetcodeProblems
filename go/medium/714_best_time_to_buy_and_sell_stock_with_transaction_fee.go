func maxProfit(prices []int, fee int) int {
    profit := 0
    cost := math.MaxInt

    for _, v := range prices {
        cost = min(cost, v - profit)
        profit = max(profit, v - cost - fee)
    }
    return profit
}
