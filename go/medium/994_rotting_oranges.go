func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    q := [][]int{}
    freshCount := 0
    for r := range grid {
        for c := range grid[0] {
            if grid[r][c] == 2 {
                q = append(q, []int{r, c})
            }
            if grid[r][c] == 1 {
                freshCount++
            }
        }
    }
    minuteCount := 0
    for len(q) > 0 && freshCount > 0 {
        size := len(q)
        for range size {
            now := q[0]
            q = q[1:]
            for _, dir := range [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
                if r, c := now[0]+dir[0], now[1]+dir[1]; r >= 0 && r < rows && c >= 0 && c < cols {
                    if grid[r][c] == 1 {
                        grid[r][c] = 2
                        freshCount--
                        q = append(q, []int{r,c})
                    }
                }
            }
        }
        minuteCount++
    }
    if freshCount == 0 {
        return minuteCount
    }
    return -1
}
