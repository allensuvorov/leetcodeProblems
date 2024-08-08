func generateMatrix(n int) [][]int {
	left, top := 0, 0
    	right, bottom := n-1, n-1
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for count := 1; count <= n*n; {
		for row, col := top, left; col <= right; col++ {
			ans[row][col] = count
			count++
		}
		top++
		for row, col := top, right; row <= bottom; row++ {
			ans[row][col] = count
			count++
		}
		right--
		for row, col := bottom, right; col >= left; col-- {
			ans[row][col] = count
			count++
		}
		bottom--
		for row, col := bottom, left; row >= top; row-- {
			ans[row][col] = count
			count++
		}
		left++
	}
	return ans
}
