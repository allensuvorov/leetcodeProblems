func pivotIndex(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    left, right := 0, sum
    for i, v := range nums {
        right -= v
        if i > 0 {
            left += nums[i-1]
        }
        if left == right {
            return i
        }
    }
    return -1
}
