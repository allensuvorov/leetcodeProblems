func findMaxAverage(nums []int, k int) float64 {
    sum, maxSum := 0, math.MinInt
    for i, v := range nums {
        sum += v
        if i - (k - 1) >= 0 {
            maxSum = max(maxSum, sum)
            sum -= nums[i - (k - 1)]
        }
    }
    return float64(maxSum)/float64(k)
}
