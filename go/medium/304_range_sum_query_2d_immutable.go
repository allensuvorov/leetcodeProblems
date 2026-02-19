type NumMatrix struct {
    prefSum [][]int // 2d prefix sum top-left to bottom-wright
}


func Constructor(matrix [][]int) NumMatrix {
    prefSum := get2DPrefSum(matrix)
    for r := range len(prefSum) {
        fmt.Println(prefSum[r])
    }
    return NumMatrix{prefSum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    res := this.prefSum[row2][col2]
    var leftDeductable, topDeductable, crossDeductable int
    if col1 > 0 {
        leftDeductable = this.prefSum[row2][col1-1]
    }
    if row1 > 0 {
        topDeductable = this.prefSum[row1-1][col2]
    }
    if row1 > 0 && col1 > 0 {
        crossDeductable = this.prefSum[row1-1][col1-1]
    }
    return res - topDeductable - leftDeductable + crossDeductable
}

func get2DPrefSum(matrix [][]int) [][]int {
    rows, cols := len(matrix), len(matrix[0])
    res := make([][]int, rows)
    for r := range rows {
        rowPrefSum := 0
        res[r] = make([]int, cols)
        for c := range cols {
            rowPrefSum += matrix[r][c]
            res[r][c] = rowPrefSum
            if r > 0 {
                res[r][c] += + res[r-1][c]
            }
        }
    }
    return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
