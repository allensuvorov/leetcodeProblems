func findMaxAverage(nums []int, k int) float64 {
    maxSum := math.MinInt
    window := 0
    for i := range nums {
        window += nums[i]
        if i + 1 >= k {
            maxSum = max(maxSum, window)
            window -= nums[1 + i - k]
        }
    }
    return float64(maxSum)/float64(k)
}
