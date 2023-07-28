func canJump(nums []int) bool {
    r := len(nums) - 1
    for l := r - 1; l >= 0; l-- {
        if nums[l] + l >= r {
            r = l
        }
    }
    if r > 0 {
        return false
    }
    return true
}
