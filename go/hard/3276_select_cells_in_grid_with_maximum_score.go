// backtracking
func maxScore(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0]) 

    for r := range grid {
        slices.SortFunc(grid[r], func(a, b int)int{ return b - a })
    }

    // backtracking
    var dfs func(usedRows, last int) int
    dfs = func(usedRows, last int) int {
        best := 0
        for r := range rows {
            if usedRows & (1 << r) == 0 {
                for c := range cols {
                    if grid[r][c] < last {
                        score := dfs(usedRows|(1<<r), grid[r][c]) + grid[r][c]
                        best = max(best, score)
                        break
                    }
                }
            }   
        }
        return best
    }

    return dfs(0, 101)
}

// bit manipulation TODO: study
func maxScore(grid [][]int) int {
    pos := make(map[int]int)
    for i, row := range grid {
        for _, x := range row {
            pos[x] |= 1 << i
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    f := make([]int, 1 << len(grid))
    for x, mask := range pos {
        for j := range f {
            for t, lb := mask, 0; t > 0; t ^= lb {
                lb = t & -t    
                if j & lb == 0 {
                    f[j] = max(f[j], f[j|lb] + x)
                }
            }
        }
    }
    return f[0]
}
