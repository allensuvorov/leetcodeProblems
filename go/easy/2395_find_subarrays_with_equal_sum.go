func findSubarrays(nums []int) bool {
    set := make(map[int]bool)
    for i := 0; i < len(nums) - 1; i++ {
        if set[nums[i] + nums[i+1]] {
            return true
        } else {
            set[nums[i] + nums[i+1]] = true
        }
    }
    return false
}
