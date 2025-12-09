func orangesRotting(grid [][]int) int {
    freshCount := 0
    rows := len(grid)
    cols := len(grid[0])
    q := [][]int{}
    for r := range rows {
        for c := range cols {
            if grid[r][c] == 1 {
                freshCount++
            }
            if grid[r][c] == 2 {
                q = append(q, []int{r, c})
            }
        }
    }
    res := 0
    for len(q) > 0 && freshCount > 0 {
        size := len(q)
        for range size {
            now := q[0]
            q = q[1:]
            for _, d := range [][]int{{1,0},{-1,0},{0,1},{0,-1}} {
                nr, nc := now[0] + d[0], now[1] + d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
                    if grid[nr][nc] == 1 {
                        grid[nr][nc] = 2
                        freshCount--
                        q = append(q, []int{nr, nc})
                    }
                }
            }
        }
        res++
    }

    if freshCount == 0 {
        return res
    }
    return -1
}
