func findPeakElement(nums []int) int {
    size := len(nums)
    if size == 1 {
        return 0
    }
    if nums[0] > nums[1] {
        return 0
    }
    if nums[size - 1] > nums[size - 2] {
        return size - 1
    }

    l := 1
    r := len(nums) - 2
    for l <= r {
        m := l + (r - l)/2
        switch {
        case nums[m] > nums[m-1] && nums[m] > nums[m+1]:
            return m
        case nums[m] > nums[m-1]:
            l = m + 1
        case nums[m] > nums[m+1]:
            r = m - 1
        default:
            l = m + 1
        }
    }
    return -1
}
