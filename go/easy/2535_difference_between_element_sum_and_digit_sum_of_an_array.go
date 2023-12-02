func differenceOfSum(nums []int) int {
    elmSum, dgtSum := 0, 0
    for _, v := range nums {
        elmSum += v
        for v > 0 {
            dgtSum += v % 10
            v /= 10
        }
    }
    return elmSum - dgtSum
}
