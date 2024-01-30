func maxSubArray(nums []int) int {
    curSum := 0
    curMax := nums[0]
    for _, v := range nums {
        curSum = max(v, v + curSum) // is that past helping?
        curMax = max(curMax, curSum)
    }
    return curMax
}
