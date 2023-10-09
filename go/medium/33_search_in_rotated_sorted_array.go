func search(nums []int, target int) int {
    if nums[0] <= nums[len(nums)-1] {
        return binarySearchTarget(nums, target)
    } else {
        pivotIndex := binarySearchPivotIndex(nums)
        if target >= nums[0] && target <= nums[pivotIndex] {
            return binarySearchTarget(nums[:pivotIndex+1], target)
        }
        if target >= nums[pivotIndex+1] && target <= nums[len(nums)-1] {
            res := binarySearchTarget(nums[pivotIndex+1:], target)
            if res == -1 {
                return res
            }
            return res + pivotIndex + 1
        }
    }
    return -1
}

func binarySearchTarget(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l <= r{
        m := l + (r - l) / 2
        guess := nums[m]
        switch {
        case guess == target:
            return m
        case guess < target:
            l = m + 1
        case guess > target:
            r = m - 1
        }
    }
    return - 1
}

func binarySearchPivotIndex(nums []int) int {
    end := len(nums) - 1
    l := 0
    r := end
    for l < r {
        m := l + (r - l + 1)/2
        if nums[m] > nums[end] {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}
