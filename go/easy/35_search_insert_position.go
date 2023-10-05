func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        m := (l + r)/2
        guess := nums[m]
        switch {
        case target == guess:
            return m
        case target > guess:
            l = m + 1
        case target < guess:
            r = m - 1
        }
    }
    return l
}
