func sortedSquares(nums []int) []int {
    size := len(nums)
    res := make([]int, size, size)
    l := 0
    r := size - 1
    top := size - 1
    for l <= r {
        ll := nums[l]*nums[l]
        rr := nums[r]*nums[r]
        if ll > rr {
            l++
            res[top] = ll
        } else {
            r--
            res[top] = rr
        }
        top--
    }
    return res
}

// compute squires once, two pointers is the same
func sortedSquares(nums []int) []int {
    for i := range nums {
        nums[i] = nums[i] * nums[i]
    }
    l, r := 0, len(nums) - 1
    result := make([]int, len(nums))
    for i := len(result) - 1; i >= 0; i-- {
        if nums[l] < nums[r] {
            result[i] = nums[r]
            r--
        } else {
            result[i] = nums[l]
            l++
        }
    }
    return result
}
