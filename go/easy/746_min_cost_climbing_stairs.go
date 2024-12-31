func minCostClimbingStairs(cost []int) int {
    a, b := 0, 0
    for _, v := range cost {
        a, b = b, v + min(a, b)
    }
    return min(a, b)
}
