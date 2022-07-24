func containsDuplicateFast(nums []int) bool {
	uniqs := make(map[int]bool)
	for _, num := range nums {
		if _, contains := uniqs[num]; contains {
			return true
		}
		uniqs[num] = true
	}
	return false
}

func containsDuplicateSlow(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}