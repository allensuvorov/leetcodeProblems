func spiralOrder(matrix [][]int) []int {
    res := []int{}
    t, l, b, r := 0, 0, len(matrix) - 1, len(matrix[0]) - 1
    rows, cols := len(matrix), len(matrix[0])
    
    for len(res) < rows*cols {
        for i, j := t, l; j <= r && len(res) < rows*cols; j++ {
            res = append(res, matrix[i][j])
        }
        t++
        for i, j := t, r; i <= b && len(res) < rows*cols; i++ {
            res = append(res, matrix[i][j])
        }
        r--
        for i, j := b, r; j >= l && len(res) < rows*cols; j-- {
            res = append(res, matrix[i][j])
        }
        b--
        for i, j := b, l; i >= t && len(res) < rows*cols; i-- {
            res = append(res, matrix[i][j])
        }
        l++
    }
    return res   
}
