func generateMatrix(n int) [][]int {
	l, t := 0, 0
    r, b := n-1, n-1
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for count := 1; count <= n*n; {
		for row, col := t, l; col <= r; col++ {
			ans[row][col] = count
			count++
		}
		t++
		for row, col := t, r; row <= b; row++ {
			ans[row][col] = count
			count++
		}
		r--
		for row, col := b, r; col >= l; col-- {
			ans[row][col] = count
			count++
		}
		b--
		for row, col := b, l; row >= t; row-- {
			ans[row][col] = count
			count++
		}
		l++
	}
	return ans
}
