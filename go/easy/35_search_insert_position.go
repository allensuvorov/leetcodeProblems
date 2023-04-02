func searchInsert(nums []int, target int) int {
    l := 0
    r := len(nums)-1
    for l <= r {
        m := (l+r)/2
        switch {
        case nums[m] == target: return m
        case nums[m] < target: l = m + 1
        case nums[m] > target: r = m - 1
        }
    }
    return l
}
