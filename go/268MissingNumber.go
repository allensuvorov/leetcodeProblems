func missingNumber(nums []int) int {
	var sum int = len(nums) * (len(nums) + 1) / 2
	for i := range nums {
		sum -= nums[i]
	}
	return sum
}