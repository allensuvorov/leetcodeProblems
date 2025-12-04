// O(n) time
func maxSubArray(nums []int) int {
    curSum := 0
    curMax := nums[0]
    for _, v := range nums {
        curSum = max(v, v + curSum) // is that past helping?
        curMax = max(curMax, curSum)
    }
    return curMax
}

// O(n^2) time
func maxSubArray(nums []int) int {
    ans := math.MinInt
    for i := range nums {
        curSum := 0
        for j := i; j < len(nums); j++ {
            curSum += nums[j]
            ans = max(ans, curSum)
        }
    }
    return ans
}
