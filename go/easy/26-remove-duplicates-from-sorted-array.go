func removeDuplicates(nums []int) int {
    lastUniq := 0
    for i := 1; i < len(nums); i ++ {
        if nums[i] != nums[lastUniq] {
            lastUniq++
            nums[lastUniq] = nums[i]
        }
    }
    return lastUniq + 1
}
