func minCostClimbingStairs(cost []int) int {
    prev, curr := 0, 0
    for _, v := range cost {
        prev, curr = curr, v + min(curr, prev)
    }
    return min(prev, curr)
}
