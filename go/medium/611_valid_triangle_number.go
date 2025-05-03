func triangleNumber(nums []int) int {
    // n^2
    n := len(nums)
    slices.Sort(nums)

    res := 0
    for a := range nums {
        for b, c := a + 1, a + 2; nums[a] > 0 && b < n - 1; b++ {
            for c < n && nums[a] + nums[b] > nums[c] {
                c++
            }
            res += c - b - 1
        }
    }
    return res
}
