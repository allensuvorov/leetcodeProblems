func findMiddleIndex(nums []int) int {
    lSum, rSum := 0, 0

    for i := range nums {
        rSum += nums[i]
    }

    for i := range nums {
        rSum -= nums[i]
        if lSum == rSum {
            return i
        }
        lSum += nums[i]
    }
    return -1
}
