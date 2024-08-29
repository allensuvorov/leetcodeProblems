func maxSubarraySumCircular(nums []int) int {
	curMax, curMin := 0, 0
	maxSum, minSum := nums[0], nums[0]
	totalSum := 0

	for _, num := range nums {
		// Normal Kadane's
		curMax = max(curMax, 0) + num
		maxSum = max(maxSum, curMax)

		// Kadane's but with min to find minimum subarray
		curMin = min(curMin, 0) + num
		minSum = min(minSum, curMin)

		totalSum += num
	}

	if totalSum == minSum {
		return maxSum
	}

	return max(maxSum, totalSum-minSum)
}
