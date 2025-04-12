func maximumCount(nums []int) int {
    n := len(nums)
    l, r := 0, n - 1
    for l <= r {
        m := l + (r - l)/2
        if nums[m] < 0 {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    neg := r + 1

    l, r = 0, n - 1
    for l <= r {
        m := l + (r - l)/2
        if nums[m] > 0 {
            r = m - 1
        } else {
            l = m + 1
        }
    }
    pos := n - l

    return max(neg, pos)
}
