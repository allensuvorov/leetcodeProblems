func isMonotonic(nums []int) bool {
    inc := true
    dec := true
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i + 1] {
            inc = false
        }
        if nums[i] < nums[i + 1] {
            dec = false
        }
    }
    return dec || inc
}
