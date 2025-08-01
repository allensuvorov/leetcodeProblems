func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    for i := range res {
        row := make([]int, i+1)
        for j := range row {
            if j == 0 || j == i {
                row[j] = 1
            } else if i > 1 {
                row[j] = res[i-1][j-1] + res[i-1][j]
            }
        }
        res[i] = row
    }
    return res
}
