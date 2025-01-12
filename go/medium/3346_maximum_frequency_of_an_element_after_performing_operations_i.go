func maxFrequency(nums []int, k int, numOperations int) int {
    slices.Sort(nums) // sorting
    counts := make(map[int]int) // frequency map
    for _, v := range nums {
        counts[v]++
    }
    
    res := 0
    
    // some unchanged
    for i, l, r := 0, 0, 0; i < len(nums); i++ {
        // keep left pointer within nums[i] - k
        for l <= i && nums[l] < nums[i] - k {
            l++
        }
        // keep right pointer within nums[i] + k
        for r < len(nums) && nums[r] <= nums[i] + k{
            spread := r - l + 1
            res = max(res, min(spread, counts[nums[i]] + numOperations))
            r++
        }
        r--
    }
    // change everything
    for l, r := 0, 0; r < len(nums); r++ {
        for l <= r && nums[l] < nums[r] - 2 * k {
            l++
        }
        spread := r - l + 1
        if spread <= numOperations {
            res = max(res, spread)
        }
    }
    return res
}
