func countMatchingSubarrays(nums []int, pattern []int) int {
    n := len(nums)
    m := len(pattern)
    res := 0
    for i := range nums {
        match := true
        k := 0
        for ; k < m && i < n - m; k++ {
            c1 := pattern[k] == 1 && nums[i + k + 1] > nums[i + k]
            c2 := pattern[k] == 0 && nums[i + k + 1] == nums[i + k]
            c3 := pattern[k] == -1 && nums[i + k + 1] < nums[i + k]
            if !(c1 || c2 || c3) {
                match = false
            }
        }

        if match && k == m{
            res++
        }
    }
    return res
}
