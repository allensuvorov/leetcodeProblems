func findMaxAverage(nums []int, k int) float64 {
    maxSum := math.MinInt
    curSum := 0
    for i := range nums {
        curSum += nums[i]
        if i + 1 >= k { // full window
            maxSum = max(maxSum, curSum)
            curSum -= nums[i + 1 - k] // cut tail
        }
    }
    return float64(maxSum)/float64(k)
}
