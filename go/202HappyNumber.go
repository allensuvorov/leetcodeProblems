func isHappy(n int) bool {
	var d int
	var s int

	for {
		d = n % 10
		s += d * d
		n /= 10
		if n == 0 {
			n = s
			if s == 1 {
				return true
			}
			s = 0
		}
	}
}