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

// linear extra space
func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefixProduct := make([]int, n)
    suffixProduct := make([]int, n)

    prefixProduct[0] = 1
    for i := 1; i < n; i++ {
        prefixProduct[i] = prefixProduct[i-1] * nums[i-1]
    }

    suffixProduct[n-1] = 1
    for i := n - 2; i >= 0; i-- {
        suffixProduct[i] = suffixProduct[i+1] * nums[i+1]
    }

    result := make([]int, n)
    for i := range n {
        result[i] = prefixProduct[i] * suffixProduct[i]
    }
    return result
}
