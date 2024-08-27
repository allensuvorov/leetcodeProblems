func longestContinuousSubstring(s string) int {
	res := 1
	cur := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i+1]-s[i] == 1 {
			cur++
		} else {
			res = max(res, cur)
			cur = 1
		}
	}
	res = max(res, cur)
	return res
}
