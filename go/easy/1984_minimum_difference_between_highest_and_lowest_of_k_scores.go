func minimumDifference(nums []int, k int) int {
    // sort, then two pointers - n log n
    sort.Ints(nums)
    min := 100001
    for i := 0; i <= len(nums) - k; i++ {
        dif := nums[i+k-1] - nums[i]
        if dif < min {
            min = dif
        }
    }
    return min
}
