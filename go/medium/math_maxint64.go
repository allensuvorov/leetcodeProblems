func gridGame(grid [][]int) int64 {
    n := len(grid[0])
    
    rowSum := make([]int64, 2)

    for i := range n {
        rowSum[0] += int64(grid[0][i])
        rowSum[1] += int64(grid[1][i])
    }
    
    res := int64(math.MaxInt64)

    rowRem := make([]int64, 2)
    rowRem[0] = rowSum[0]

    for i := range n {
        rowRem[0] -= int64(grid[0][i])
        if i > 0 {
            rowRem[1] += int64(grid[1][i-1])
        }
        res = min(res, max(rowRem[0], rowRem[1]))
    }

    return res
}
