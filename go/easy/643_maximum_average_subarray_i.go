func findMaxAverage(nums []int, k int) float64 {
    maxAverage := float64(math.MinInt)
    sum := 0
    for i, v := range nums {
        sum += v
        if i - (k - 1) >= 0 {
            maxAverage = math.Max(maxAverage, float64(sum) / float64(k))
            sum -= nums[i - (k - 1)]
        }
    }
    return maxAverage
}
