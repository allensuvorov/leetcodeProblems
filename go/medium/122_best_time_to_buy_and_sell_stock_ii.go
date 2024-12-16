func maxProfit(prices []int) int {
    maxProf := 0
    for sellDay := 1; sellDay < len(prices); sellDay++ {
        buyDay := sellDay - 1
        profit := prices[sellDay] - prices[buyDay]
        if profit > 0 {
            maxProf += profit
        }
    }
    return maxProf
}
