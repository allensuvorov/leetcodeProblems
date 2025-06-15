func countMatchingSubarrays(nums []int, pattern []int) int {
    n := len(nums)
    m := len(pattern)
    res := 0
    for i := range nums {
        match := true
        k := 0
        for ; k < m && i < n - m; k++ {
            inc := pattern[k] == 1 && nums[i + k + 1] > nums[i + k]
            non := pattern[k] == 0 && nums[i + k + 1] == nums[i + k]
            dec := pattern[k] == -1 && nums[i + k + 1] < nums[i + k]
            if !inc && !non && !dec {
                match = false
            }
        }

        if match && k == m{
            res++
        }
    }
    return res
}
