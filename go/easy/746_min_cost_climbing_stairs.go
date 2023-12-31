func minCostClimbingStairs(cost []int) int {
    a, b := 0, 0
    for _, v := range cost {
        a, b = b, min(a, b) + v
    }
    return min(a, b) 
}
