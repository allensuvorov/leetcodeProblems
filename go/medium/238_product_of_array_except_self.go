func productExceptSelf(nums []int) []int {
    
    left := make([]int, len(nums))
    right := make([]int, len(nums))
    
    product := 1
    for i, v := range nums {
        product *= v
        left[i] = product
    }

    product = 1
    for i := len(nums) - 1; i >= 0; i-- {
        product *= nums[i]
        right[i] = product
    }

    res := make([]int, len(nums))
    res[0] = right[1]
    res[len(res)-1] = left[len(res)-2]
    for i := 1; i < len(res)-1; i++ {
        res[i] = left[i-1] * right[i+1]
    }
    return res
}
