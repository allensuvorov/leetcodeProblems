func maxProfit(prices []int) int {
    ans, profit, maxProfit := 0, 0, 0
    minPrice := prices[0]
    for _, v := range prices {
        if v < minPrice {
            minPrice = v
        }
        profit = v - minPrice
        if maxProfit <= profit {
            maxProfit = profit
        } else {
            ans += maxProfit
            maxProfit = 0
            minPrice = v
        }
    }
    ans += maxProfit
    return ans
}
