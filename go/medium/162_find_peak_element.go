func findPeakElement(nums []int) int {
    l := 1
    r := len(nums) - 1
    for l < r {
        m := l + (r - l)/2
        if nums[m] > nums[m+1]{
             r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func findPeakElement(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		m := l + (r-l)/2
		if m != n-1 && nums[m] < nums[m+1] {
			l = m + 1
		} else if m != 0 && nums[m] < nums[m-1] {
			r = m - 1
		} else {
			return m
		}
	}
	return l
}
