func rotate(matrix [][]int)  {
    for left := range len(matrix)/2 {
        right := len(matrix) - 1 - left
        for i := range right - left {
            r1, r2, r3, r4 := left,     left + i,     right,        right - i
            c1, c2, c3, c4 := left + i, right,        right - i,    left
            matrix[r2][c2], matrix[r3][c3], matrix[r4][c4], matrix[r1][c1] = matrix[r1][c1], matrix[r2][c2], matrix[r3][c3], matrix[r4][c4]
        }
        left++
    }
}
