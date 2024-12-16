func maxProfit(prices []int) int {
    maxProfit := 0
    minPrice := prices[0]
    for i := 1; i < len(prices); i++ {
        maxProfit = max(maxProfit, prices[i] - minPrice)
        minPrice = min(minPrice, prices[i])
    }
    return maxProfit
}
