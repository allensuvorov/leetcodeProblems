func check(nums []int) bool {
    // if there is one decrease - check last < first
    decreaseCount := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i-1] {
            decreaseCount ++
            if decreaseCount > 1 {
                return false
            }
        }
    }
    if decreaseCount == 1 && nums[0] < nums[len(nums)-1] {
        return false
    }
    return true
}
