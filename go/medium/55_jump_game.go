// r->l
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

// l->r , new high
func canJump(nums []int) bool {
    newHigh := 0
    for i, v := range nums {
        if newHigh >= i {
            newHigh = max(newHigh, i + v)
        }
    }
    return newHigh >= len(nums) - 1
}
