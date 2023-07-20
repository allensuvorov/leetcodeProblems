func removeDuplicates(nums []int) int {
    lastUniq := 1
    for i := 1; i < len(nums); i ++ {
        if nums[i] != nums[lastUniq-1] {
            nums[lastUniq] = nums[i]
            lastUniq++
        }
    }
    return lastUniq
}
