func spiralOrder(matrix [][]int) []int {
    rows, cols := len(matrix), len(matrix[0])
    size := rows*cols
    ans := make([]int, 0, size)
    
    // Initialize boundries
    t, b, l, r := 0, rows - 1, 0, cols - 1

    for len(ans) < size { 
        for i := l; i <= r && len(ans) < size; i++ {
            ans = append(ans, matrix[t][i])
        }
        t++
        for i := t; i <= b && len(ans) < size; i++ {
            ans = append(ans, matrix[i][r])
        }
        r--
        for i := r; i >= l && len(ans) < size; i-- {
            ans = append(ans, matrix[b][i])
        }
        b--
        for i := b; i >= t && len(ans) < size; i-- {
            ans = append(ans, matrix[i][l])
        }
        l++
    }
    return ans
}
