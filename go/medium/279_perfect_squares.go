var mem = map[int]int{0: 0} 

func numSquares(n int) int {
	if res, ok := mem[n]; ok {
		return res
	}
	min := n
	for i := 1; i*i <= n; i++ {
		s := numSquares(n - i*i)
		if s < min {
			min = s
		}
	}
	mem[n] = 1 + min
	return 1 + min
}
