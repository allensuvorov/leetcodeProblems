func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        m := (l + r)/2
        num := nums[m] 
        if num == target {
            return m
        }
        if num < target {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return -1
}
