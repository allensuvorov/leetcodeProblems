func minimumLines(stockPrices [][]int) int {
    sort.Slice(stockPrices, func(i, j int) bool {
        return stockPrices[i][0] < stockPrices[j][0]
    })
    
    minLines := 0
    prevSlope := math.MaxInt
    for i := 1; i < len(stockPrices); i++ {
        day, price := stockPrices[i][0], stockPrices[i][1]
        prevDay, prevPrice := stockPrices[i-1][0], stockPrices[i-1][1]
        slope := (price - prevPrice) * (day - prevDay)
        if slope != prevSlope {
            fmt.Println(slope, prevSlope)
            minLines++
        }
        prevSlope = slope
    }
    return minLines
}
