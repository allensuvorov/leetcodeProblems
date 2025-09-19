func maxSum(nums []int) int {
    // get sum of all unique non-negative values
    var sum int
    seenNonNegNums := make(map[int]bool)
    maxNum := math.MinInt

    for _, v := range nums {
        if v >= 0 && !seenNonNegNums[v] {
            seenNonNegNums[v] = true
            sum += v
        }
        maxNum = max(v, maxNum)
    }

    // if only negative values - return the largest one
    if len(seenNonNegNums) == 0 {
        return maxNum
    }
    return sum
}
