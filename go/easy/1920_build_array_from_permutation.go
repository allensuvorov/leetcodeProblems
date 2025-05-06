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

// Do as required
func buildArray(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    for i, v := range nums {
        ans[i] = nums[v]
    }
    return ans
}
