func maximumUniqueSubarray(nums []int) int {
    windowSet := make([]bool, 10_001)
    windowSum := 0
    l := 0
    res := 0
    for _, v := range nums {
        windowSum += v
        // sweep the tail to remove all dups from window
        for windowSet[v] {
            windowSum -= nums[l]            
            windowSet[nums[l]] = false 
            l++
        }
        windowSet[v] = true

        res = max(res, windowSum)
    }
    return res
}
