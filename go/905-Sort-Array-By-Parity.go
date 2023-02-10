func sortArrayByParity(nums []int) []int {
    l := 0
    r := len(nums)-1
    for l < r {
        if nums[l]%2 == 0 {
            l++
            continue
        }
        if nums[r]%2 != 0 {
            r--
            continue
        }
        nums[l], nums[r] = nums[r], nums[l]
    }
    return nums
}
