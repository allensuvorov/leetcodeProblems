func findMaxConsecutiveOnes(nums []int) int {
    res := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            count++
            res = max(res, count)
        } else {
            count = 0
        }
    }
    return res
}
