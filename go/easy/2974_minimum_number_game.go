func numberGame(nums []int) []int {
	fa := make([]int, 101)
	for _, v := range nums {
		fa[v]++
	}

	i := 0
    cur := 0
	for i < len(nums) && cur < len(fa) {
		if fa[cur] == 0 {
			cur++
            continue
		}
        nums[i] = cur
        fa[cur]--
        i++
	}

    for i := 1; i < len(nums); i += 2 {
        nums[i], nums[i-1] = nums[i-1], nums[i]
    }
    return nums
}
