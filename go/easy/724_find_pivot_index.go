func pivotIndex(nums []int) int {
    totalSum := 0
    for _, v := range nums{
        totalSum += v
    }

    leftSum := 0
    rightSum := totalSum

    for i, v := range nums{
        rightSum -= v
        if leftSum == rightSum {
            return i
        }
        leftSum += v
    }

    return -1
}

