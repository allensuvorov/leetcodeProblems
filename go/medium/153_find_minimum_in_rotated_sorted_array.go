func findMin(nums []int) int {
    // find the first thats less then the last
    size := len(nums)
    l := 0
    r := size - 1
    for l < r {
        m := l + (r - l)/2
        if nums[m] < nums[size - 1] {
            r = m
        } else {
            l = m + 1
        }
    }
    return nums[l]
}
