func findMaxAverage(nums []int, k int) float64 {
    maxSum := math.MinInt
    curSum := 0
    for i := range nums {
        curSum += nums[i]
        if i >= k { // we have a tail item - cut it
            curSum -= nums[i - k]
        }
        if i + 1 >= k { // we have a k lenth subarray
            maxSum = max(maxSum, curSum)
        }
    }
    return float64(maxSum)/float64(k)
}
