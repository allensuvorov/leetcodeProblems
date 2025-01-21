func gridGame(grid [][]int) int64 {
    n := len(grid[0])
    
    var rowZeroSum int64

    for i := range n {
        rowZeroSum += int64(grid[0][i])
    }
    
    rowRem := []int64 {rowZeroSum, 0}
    res := int64(math.MaxInt64)

    for i := range n {
        rowRem[0] -= int64(grid[0][i])
        if i > 0 {
            rowRem[1] += int64(grid[1][i-1])
        }
        res = min(res, max(rowRem[0], rowRem[1]))
    }

    return res
}
