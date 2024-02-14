func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, cols := len(image), len(image[0])
    canvas := image[sr][sc]
    if canvas == color {
	    return image
	}
    stack := [][]int{{sr, sc}}
    for len(stack) > 0 {
        r := stack[len(stack) - 1][0]
        c := stack[len(stack) - 1][1]
        stack = stack[:len(stack) - 1]
        if image[r][c] == canvas {
			image[r][c] = color
			if r > 0 {
				stack = append(stack, []int{r-1, c})
			}
			if r < rows-1 {
				stack = append(stack, []int{r+1, c})
			}
			if c > 0 {
				stack = append(stack, []int{r, c-1})
			}
			if c < cols-1 {
				stack = append(stack, []int{r, c+1})
			}
		}
    }
	return image
}
