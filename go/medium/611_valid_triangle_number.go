func triangleNumber(nums []int) int {
    // n^2 * log n
    n := len(nums)
    slices.Sort(nums)
    binSearch := func(b int, target int) int {
        l, r := b+1, n-1
        c := b
        for l <= r {
            m := l + (r - l)/2
            if nums[m] >= target {
                r = m - 1
            } else {
                c = m
                l = m + 1
            }
        }
        return c - b
    }

    res := 0
    for a := range nums {
        for b := a + 1; nums[a] > 0 && b < n - 1; b++ {
            res += binSearch(b, nums[a] + nums[b])
        }
    }
    return res
}
