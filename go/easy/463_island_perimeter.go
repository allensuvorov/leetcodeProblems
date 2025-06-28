func islandPerimeter(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    res := 0
    for r := range rows {
        for c := range cols {
            if grid[r][c] == 1 {
                // check neighbors
                if r == 0 || grid[r-1][c] == 0 {
                    res++
                }
                if r == rows-1 || grid[r+1][c] == 0 {
                    res++
                }
                if c == 0 || grid[r][c-1] == 0 {
                    res++
                }
                if c == cols-1 || grid[r][c+1] == 0 {
                    res++
                }
            }
        }
    }
    return res
}
