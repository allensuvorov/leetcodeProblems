func imageSmoother(img [][]int) [][]int {
    rows := len(img)
    cols := len(img[0])
    directions := [][]int{
        {-1,-1},{-1,0},{-1,1},
        {0,-1},{0,1},
        {1,-1},{1,0},{1,1},
    }
    res := make([][]int, rows)
    for r := range rows {
        res[r] = make([]int, cols)
    }

    for r := range rows {
        for c := range cols {
            sum := img[r][c]
            count := 1
            for _, direction := range directions {
                neiRow := r + direction[0]
                neiCol := c + direction[1]
                if neiRow >= 0 && neiRow < rows && neiCol >= 0 && neiCol < cols {
                    sum += img[neiRow][neiCol]
                    count++
                }
            }
            res[r][c] = sum/count
        }
    }
    return res
}
