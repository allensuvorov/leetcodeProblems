func maximumSumOfHeights(heights []int) int64 {
    n := len(heights)
    var maxSum int64
    
    for i := range n {
        sum := int64(heights[i])
        minSoFar := heights[i]
        for r := i + 1; r < n; r++ {
            minSoFar = min(minSoFar, heights[r])
            sum += int64(minSoFar)
        }
        minSoFar = heights[i]
        for l := i - 1; l >= 0; l-- {
            minSoFar = min(minSoFar, heights[l])
            sum += int64(minSoFar)
        }
        maxSum = max(maxSum, sum)
        fmt.Println(maxSum)
    }
    return maxSum
}
