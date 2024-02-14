func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, cols := len(image), len(image[0])
    canvas := image[sr][sc]
    if canvas == color {
	    return image
	}
    stack := [][]int{{sr, sc}}
    for len(stack) > 0 {
        r, c := stack[len(stack) - 1][0], stack[len(stack) - 1][1]
        stack = stack[:len(stack) - 1]
        image[r][c] = color
        for _, pos := range [][]int{{1,0},{-1,0},{0,1},{0,-1}} {
            nr, nc := r+pos[0], c+pos[1]
            if nr >= 0 && nr < rows && nc >= 0 && nc < cols && image[nr][nc] == canvas {
                stack = append(stack, []int{nr, nc})
            }
        }
		
    }
	return image
}
