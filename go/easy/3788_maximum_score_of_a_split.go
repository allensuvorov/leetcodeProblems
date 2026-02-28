func maximumScore(nums []int) int64 {
    prefSum := make([]int, len(nums))
    curSum := 0
    for i, v := range nums {
        curSum += v
        prefSum[i] = curSum
    }

    var res int64 = math.MinInt64
    sufMin := nums[len(nums)-1]
    for i := len(nums) - 2; i >= 0; i-- {
        score := int64(prefSum[i]) - int64(sufMin)
        res = max(res, score)
        sufMin = min(sufMin, nums[i])
    }
    return res
}
