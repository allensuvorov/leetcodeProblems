func longestOnes(nums []int, k int) int {
	zeros := 0
    if nums[0] == 0 {
        zeros = 1
    }
	l, r := 0, 0
	
    res := 1
    if k == 0 &&  nums[0] == 0 {
        res = 0
    }

	for l <= r && r < len(nums) {

		if (l == r || zeros <= k) && r < len(nums)-1 {
			r++
			if nums[r] == 0 {
				zeros++
			}
		} else {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		if zeros <= k {
			res = max(res, r-l+1)
			fmt.Println(res)
		}
	}
	return res
}
