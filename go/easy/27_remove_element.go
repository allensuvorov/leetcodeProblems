func removeElement(nums []int, val int) int {
    l := 0
    r := len(nums) - 1
    for l <= r {
        if nums[l] != val {
            l++
            continue
        }
        if nums[r] == val {
            r--
            continue
        }
        nums[l], nums[r] = nums[r], nums[l]
    }
    return l
}
