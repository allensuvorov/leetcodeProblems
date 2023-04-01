func finalPrices(prices []int) []int {
    for i := range prices {
        for j:= i+1; j < len(prices); j++{
            if prices[j] <= prices[i]{
                prices[i] -= prices[j]
                break
            }
        } 
    }
    return prices
}
