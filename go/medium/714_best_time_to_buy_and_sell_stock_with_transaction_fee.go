func maxProfit(prices []int, fee int) int {
	cost := math.MaxInt
	profit := 0

	for _, price := range prices {
        cost = min(cost, price - profit)
        profit = max(profit, price - cost - fee)
	}

    return profit
}
