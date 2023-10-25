func maxSubsequence(nums []int, k int) []int {
    // solution 1 - O(n * log n)
    // Sort, get frequency map
    // Collect from source
    size := len(nums)
    sorted := make([]int, size)
    copy(sorted, nums)
    sort.Ints(sorted)
    
    fm := map[int]int{}
    
    for i := size - k; i < size; i ++ {
        fm[sorted[i]]++
    }

    res := make([]int, 0, k)

    for _, v := range nums {
        if fm[v] > 0 {
            res = append(res, v)
            fm[v]--
        }
    }
    return res
}
