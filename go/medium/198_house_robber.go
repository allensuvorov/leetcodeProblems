func rob(nums []int) int {
    curMax, prevMax := 0, 0
    for _, v := range nums {
        curMax, prevMax = max(prevMax + v, curMax), curMax
    }
    return curMax
}
