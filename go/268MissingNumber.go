func missingNumber(nums []int) int {
	n := len(nums)
	var sum int
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	// fmt.Println (sum, n*(n+1)/2)
	return n*(n+1)/2 - sum
	// triangle area = n*n/2
}