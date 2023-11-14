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
