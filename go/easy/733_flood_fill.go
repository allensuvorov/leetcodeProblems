func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, cols := len(image), len(image[0])
	curColor := image[sr][sc]
	newColor := color
	if curColor == newColor {
		return image
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if image[r][c] == curColor {
			image[r][c] = newColor
			if r > 0 {
				dfs(r-1, c)
			}
			if r < rows-1 {
				dfs(r+1, c)
			}
			if c > 0 {
				dfs(r, c-1)
			}
			if c < cols-1 {
				dfs(r, c+1)
			}
		}
	}
	dfs(sr, sc)
	return image
}
