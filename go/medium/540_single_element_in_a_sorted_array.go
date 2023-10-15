func singleNonDuplicate(nums []int) int {
    l := 0
    r := len(nums) - 1
    for l < r {
        m := l + (r - l)/2
        if nums[m] == nums[m + 1] { 
            if (r - m + 1) % 2 == 0 {
                r = m - 1 // pop right
            } else {
                l = m // pop left
            }
        } else if nums[m] == nums[m - 1] {
            if (m - l + 1) % 2 == 0 {
                l = m + 1 // pop left
            } else {
                r = m // pop right
            }
        } else {
            return nums[m]
        }
    }
    return nums[l]
}
