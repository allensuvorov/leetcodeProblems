func pivotIndex(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    l, r := 0, sum
    for i, v := range nums {
        if i > 0 {
            l += nums[i-1]
        }
        r -= v
        if l == r {
            return i
        }
    }
    return -1
}
