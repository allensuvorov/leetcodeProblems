func buildArray(nums []int) []int {
    q := len(nums)
    for i, v := range nums {
        nums[i] += q * (nums[v] % q)
    }
    for i := range nums {
        nums[i] /= q
    }

    return nums
}
