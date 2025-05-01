func searchRange(nums []int, target int) []int {
    start := -1 // bin search to find starting position
    l, r := 0, len(nums) - 1
    for l <= r {
        m := l + (r - l)/2
        if nums[m] == target {
            start = m
            r = m - 1
        } else if nums[m] > target {
            r = m - 1
        } else if nums[m] < target {
            l = m + 1
        }
    }

    end := -1 // bin search to find ending position
    l, r = 0, len(nums) - 1
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
