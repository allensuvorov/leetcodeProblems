func minimumArea(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    minRow, maxRow := rows, -1
    minCol, maxCol := cols, -1
    for r := range rows {
        for c := range cols {
            if grid[r][c] == 1 {
               minRow = min(minRow, r) 
               maxRow = max(maxRow, r) 
               minCol = min(minCol, c) 
               maxCol = max(maxCol, c) 
            }
        }
    }
    return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
