func maximumDifference(nums []int) int {
    minVal := math.MaxInt
    maxDiff := -1
    for _, v := range nums {
        minVal = min(minVal, v)
        if v > minVal {
            maxDiff = max(maxDiff, v - minVal)
        }
    }
    return maxDiff
}
