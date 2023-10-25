func maxSubsequence(nums []int, k int) []int {
    // solution - sort
    // time O(n * log n)
    size := len(nums)

    // Sort
    sorted := make([]int, size)
    copy(sorted, nums)
    sort.Ints(sorted)
    
    // Frequency map of top range
    fm := map[int]int{}
    for i := size - k; i < size; i ++ {
        fm[sorted[i]]++
    }

    // Collect from source
    res := make([]int, 0, k)
    for _, v := range nums {
        if fm[v] > 0 {
            res = append(res, v)
            fm[v]--
        }
    }
    return res
}
