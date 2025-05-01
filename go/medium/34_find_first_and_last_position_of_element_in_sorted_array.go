func searchRange(nums []int, target int) []int {
    n := len(nums)
    // bin search to find starting position
    start := -1
    l, r := 0, n - 1
    for l <= r {
        m := l + (r - l)/2
        if nums[m] == target {
            r = m - 1
            start = m
        } else if nums[m] > target {
            r = m - 1
        } else if nums[m] < target {
            l = m + 1
        }
    }

    // bin search to find ending position
    end := -1
    l, r = 0, n - 1
    for l <= r {
        m := l + (r - l)/2
        if nums[m] == target {
            end = m
            l = m + 1
        } else if nums[m] > target {
            r = m - 1
        } else if nums[m] < target {
            l = m + 1
        }

    }
    return []int{start, end}
}
