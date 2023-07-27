// update min price and update min profit
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

// two pointers
func maxProfit(prices []int) int {
    maxP, l, r := 0, 0, 1 // l - buy, r - sell
    for r < len(prices); r++ {
        if prices[l] < prices[r]  { // profitable?
            profit := prices[r] - prices[l]
            if maxP < profit {
                maxP = profit
            }
        } else {
            l = r
        }
    }
    return maxP
}
