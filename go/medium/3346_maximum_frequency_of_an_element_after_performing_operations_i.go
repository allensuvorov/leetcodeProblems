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
        // keep right pointer within nums[i] + k + 1
        for r < len(nums) && nums[r] <= nums[i] + k{
            r++
        }
        window := r - l
        res = max(res, min(window, counts[nums[i]] + numOperations))
    }

    // change everything
    for l, r := 0, 0; r < len(nums); r++ {
        for l <= r && nums[l] + 2 * k < nums[r] {
            l++
        }
        window := r - l + 1
        if window <= numOperations {
            res = max(res, window)
        }
    }
    return res
}
