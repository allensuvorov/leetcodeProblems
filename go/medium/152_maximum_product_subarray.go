func maxProduct(nums []int) int {
    curMax := 1
    curMin := 1
    allMax := nums[0]
    for _, v := range nums {
        curMax, curMin = max(v, v * curMax, v * curMin), min(v, v*curMax, v*curMin)
        allMax = max(curMax, allMax)
    }
    return allMax
}
