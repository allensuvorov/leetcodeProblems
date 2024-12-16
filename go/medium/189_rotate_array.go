// O(1) extra memory
func rotate(nums []int, k int)  {
    k = k % len(nums)
    slices.Reverse(nums)
    slices.Reverse(nums[:k])
    slices.Reverse(nums[k:])
}

// O(n) extra memory
func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    a := nums[:n - k]
    b := nums[n - k:]
    c := make([]int, 0, n)
    for _, v := range b {
        c = append(c, v)
    }
    for _, v := range a {
        c = append(c, v)
    }
    copy(nums, c)
}
