func minCostClimbingStairs(cost []int) int {
    prev, curr := 0, 0
    for _, v := range cost {
        temp := curr
        curr = v + min(curr, prev)
        prev = temp
    }

    return min(prev, curr)
}
