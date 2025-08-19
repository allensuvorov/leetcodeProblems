func pivotIndex(nums []int) int {
    rightSum := 0
    for _, v := range nums {
        rightSum += v
    }
    leftSum := 0
    for i, v := range nums {
        rightSum -= v
        if leftSum == rightSum {
            return i
        }
        leftSum += v
    }
    return -1
}
