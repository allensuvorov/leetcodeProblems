func countSubarrays(nums []int, k int) int64 {
    n := len(nums)
    maxLmt := 0
    for _, v := range nums {
        maxLmt = max(maxLmt, v)
    }

    var res int64
    window := 0
    for l, r := 0, 0; l < n && r < n; r++ { // sliding window
        if nums[r] == maxLmt {
            window++
        }
        for window == k { // have a window
            if nums[l] == maxLmt { 
                window--
            }
            l++ // trim left
        }
        res += int64(l) // add left side combinations
    }
    return res
}
