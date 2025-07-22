func maximumUniqueSubarray(nums []int) int {
    windowSet := make(map[int]struct{})
    windowSum := 0
    l := 0
    res := 0
    for _, v := range nums {
        windowSum += v
        // sweep the tail to remove all dups from window
        for _, exists := windowSet[v]; exists; _, exists = windowSet[v] {
            windowSum -= nums[l]            
            delete(windowSet, nums[l]) 
            l++
        }
        windowSet[v] = struct{}{}

        res = max(res, windowSum)
    }
    return res
}
