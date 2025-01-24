func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    // right prefix products
    right := 1
    for i := n - 1; i >= 0; i-- {
        res[i] = right
        right *= nums[i]
    }
    // left prefix products
    left := 1
    for i := range nums {
        res[i] *= left
        left *= nums[i]
    }

    return res
}
