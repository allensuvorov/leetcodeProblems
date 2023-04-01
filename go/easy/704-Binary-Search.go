func search(nums []int, target int) int { 
    l := 0 
    r := len(nums)-1

    for l <= r {
        mid := (l + r)/2
        guess := nums[mid]
        switch {
            case target == guess:
                return mid
            case target < guess:
                r = mid - 1
            case target > guess:
                l = mid + 1
        }
    }
    return -1
}
