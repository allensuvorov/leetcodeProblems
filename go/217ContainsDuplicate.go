func containsDuplicate(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}