func countNegatives(grid [][]int) int {
    // O(m + n)
    // Start at top right corner
    m, n := len(grid), len(grid[0])
    i , j := 0, n - 1
    ans := 0
    for i < m && j >= 0 {
        curr := grid[i][j]
        if curr < 0 {
            ans += m - i
            j--
        } else {
            i++
        }
    } 
    return ans
}

// start a left bottom corner
func countNegatives(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    r, c := rows - 1, 0 // start a left bottom corner
    res := 0
    for r >= 0 && c < cols {
        if grid[r][c] < 0 {
            res += cols - c
            r--
        } else {
            c++
        }
    }
    return res
}
