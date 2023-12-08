func findMaxAverage(nums []int, k int) float64 {
    maxAve := float64(math.MinInt64)
    var curSum int

    for i := 0; i < k; i ++ {
        curSum += nums[i]
    }

    var curAve float64 = float64(curSum)/float64(k)
    if curAve > maxAve {
        maxAve = curAve
    }
    
    l := 0
    for r := k; r < len(nums); r++ {
        curSum -= nums[l]
        curSum += nums[r] 
        var curAve float64 = float64(curSum)/float64(k)
        if curAve > maxAve {
            maxAve = curAve
        }
        l++
    }
    return maxAve

}
