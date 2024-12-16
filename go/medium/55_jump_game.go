func canJump(nums []int) bool {
    n := len(nums)
    highestSolution := n - 1
    for i := n - 1; i >= 0; i-- {
        if i + nums[i] >= highestSolution {
            highestSolution = i
        }
    }
    return highestSolution == 0
}
