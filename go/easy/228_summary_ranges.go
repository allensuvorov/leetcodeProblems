func summaryRanges(nums []int) []string {
    start, end := 0, 0
	res := []string{}

	for i := 0; i < len(nums); i++ {
		end = i
        if i == len(nums) - 1 || nums[i+1]-nums[i] != 1 {
			rng := strconv.Itoa(nums[start])
			if start != end {
				rng += "->" + strconv.Itoa(nums[end])
			}
			res = append(res, rng)
			start = i + 1
        }
	}
	return res
}
