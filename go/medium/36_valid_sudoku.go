func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]bool
	for r := range 9 {
		for c := range 9 {
			if board[r][c] != '.' {
				i := board[r][c] - '1'
				if row[r][i] || col[c][i] || box[c/3*3+r/3][i] {
					return false
				}
				row[r][i], col[c][i], box[c/3*3+r/3][i] = true, true, true
			}
		}
	}
	return true
}
