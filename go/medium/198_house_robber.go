func rob(nums []int) int {
    // track max before prev
    curDP := 0
    prevDP := 0
    prevMaxDP := 0
    for _, v := range nums {
        curDP = v + prevMaxDP
        prevMaxDP = max(prevDP, prevMaxDP)
        prevDP = curDP
    }
    return max(prevDP, prevMaxDP)
}
