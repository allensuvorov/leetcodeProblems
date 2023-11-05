func findLengthOfLCIS(nums []int) int {
    max := 1
    cur := 1
    for i := 1; i < len(nums); i++ {
        if nums[i-1] < nums[i] {
            cur++
        } else {
            cur = 1
        }
        if cur > max {
            max = cur
        }
    }
    return max
}
