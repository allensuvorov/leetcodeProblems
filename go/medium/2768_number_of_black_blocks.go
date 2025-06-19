func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
    cells := make(map[cell]int)
    visitedBlocks := make(map[cell]bool)

    for _, v := range coordinates {
        cells[cell{v[0], v[1]}] = 1
    }

    res := make([]int64, 5)

    for _, v := range coordinates {
        r, c := v[0], v[1]
        // check block 1
        if !visitedBlocks[cell{r-1, c-1}] && r > 0 && c > 0 {
            visitedBlocks[cell{r-1, c-1}] = true
            blackCount := 1 + cells[cell{r, c-1}] + cells[cell{r-1, c-1}] + cells[cell{r-1, c}]
            res[blackCount]++
        }
        // check block 2
        if !visitedBlocks[cell{r-1, c}] && r > 0 && c < n - 1 {
            visitedBlocks[cell{r-1, c}] = true
            blackCount := 1 + cells[cell{r-1, c}] + cells[cell{r-1, c+1}] + cells[cell{r, c+1}]
            res[blackCount]++
        }
        // check block 3
        if !visitedBlocks[cell{r, c}] && r < m - 1 && c < n - 1 {
            visitedBlocks[cell{r, c}] = true
            blackCount := 1 + cells[cell{r, c+1}] + cells[cell{r+1, c+1}] + cells[cell{r+1, c}]
            res[blackCount]++
        }
        // check block 4
        if !visitedBlocks[cell{r, c-1}] && r < m - 1 && c > 0 {
            visitedBlocks[cell{r, c-1}] = true
            blackCount := 1 + cells[cell{r+1, c}] + cells[cell{r+1, c-1}] + cells[cell{r, c-1}]
            res[blackCount]++
        }
    }

    res[0] = int64((m-1) * (n-1)) - (res[1] + res[2] + res[3] + res[4])
    return res
}

type cell struct {
    x, y int
}
