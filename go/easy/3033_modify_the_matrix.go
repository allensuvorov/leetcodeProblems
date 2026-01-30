func modifiedMatrix(matrix [][]int) [][]int {
    rows, cols := len(matrix), len(matrix[0])
    answer := make([][]int, rows)
    
    for r := range rows {
        answer[r] = make([]int, cols)
        for c := range cols {
            answer[r][c] = matrix[r][c]
        }
    }

    // modify
    for c := range cols {
        maxVal := 0
        for r := range rows {
            maxVal = max(maxVal, answer[r][c])
        }
        for r := range rows {
            if answer[r][c] == -1 {
                answer[r][c] = maxVal
            }
        }
    }
    return answer
}
