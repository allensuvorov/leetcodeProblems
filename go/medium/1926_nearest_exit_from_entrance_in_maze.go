func nearestExit(maze [][]byte, entrance []int) int {
    rows, cols := len(maze), len(maze[0])
    q := [][]int{entrance}
    stepCount := 0
    for len(q) > 0 {
        rowSize := len(q)
        for range rowSize {// pick next step
            now := q[0]
            q = q[1:]
            r, c := now[0], now[1]
            if now[0] != entrance[0] || now[1] != entrance[1] { // if not at entrance
                if r == 0 || r == rows - 1 || c == 0 || c == cols - 1 { // if exit
                    return stepCount
                }
            }
            if r > 0 && maze[r - 1][c] == '.' {
                q = append(q, []int{r - 1, c})
                maze[r-1][c] = '-'
            }
            if r < rows - 1 && maze[r + 1][c] == '.' {
                q = append(q, []int{r + 1, c})
                maze[r+1][c] = '-'
            }
            if c > 0 && maze[r][c - 1] == '.' {
                q = append(q, []int{r, c - 1})
                maze[r][c-1] = '-'
            }
            if c < cols - 1 && maze[r][c + 1] == '.' {
                q = append(q, []int{r, c + 1})
                maze[r][c+1] = '-'
            }
        }
        stepCount++
    }
    return -1
}
