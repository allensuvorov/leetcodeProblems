func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    left := 1
    for i := 0; i < len(nums); i++ {
        res[i] = left
        left *= nums[i]
    }
    right := 1
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= right
        right *= nums[i]
    }
    return res
}

// second try
func productExceptSelf(nums []int) []int {
    n := len(nums)

    suff := make([]int, n)
    suff[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i-- {
        suff[i] = nums[i]*suff[i+1]
    }
    
    for i := 1; i < n; i++ {
        nums[i] *= nums[i-1]
    }

    suff[0] = suff[1]
    for i := 1; i < n - 1; i++ {
        suff[i] = nums[i-1] * suff[i+1]
    }
    suff[n-1] = nums[n-2]
    return suff
}
