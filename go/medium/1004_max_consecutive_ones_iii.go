func longestOnes(nums []int, k int) int {
	zeros := 0
	l, r := 0, 0
    res := 0

	for r < len(nums) {
        if nums[r] == 0 {
            zeros++
        }
		if zeros > k{
            if nums[l] == 0 {
				zeros--
			}
			l++
		}
		if zeros <= k {
			res = max(res, r-l+1)
		}
        r++
	}
	return res
}
