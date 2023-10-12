func searchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    if len(nums) == 0 {
        return res
    }
    res[0] = findFirstMatch(nums, target)
    res[1] = findLastMatch(nums, target)
    return res
}

func findFirstMatch(nums[]int, target int) int {
    l := 0
    r := len(nums) - 1
    for l < r {
        m := l + (r - l)/2
        if nums[m] >= target {
            r = m
        } else {
            l = m + 1
        }
    }
    if nums[l] == target {
        return l
    } else {
        return -1
    }
}

func findLastMatch(nums[]int, target int) int {
    l := 0
    r := len(nums) - 1
    for l < r {
        m := l + (r - l + 1)/2
        if nums[m] <= target {
            l = m
        } else {
            r = m - 1
        }
    }
    if nums[l] == target {
        return l
    } else {
        return -1
    }
}
