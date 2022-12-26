func maxProfit(prices []int) int {
    minPrice, profit, maxProfit := prices[0], 0, 0
    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        }
        profit = price - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}
