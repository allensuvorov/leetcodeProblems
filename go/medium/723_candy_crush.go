func candyCrush(board [][]int) [][]int {
    for flagStreaks(board) {
        crushAndDrop(board)
    }
    return board
}

func flagStreaks(board [][]int) bool {
    rows := len(board)
    cols := len(board[0])
    streakExists := false
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            // if the same switch sign
            if board[r][c] != 0 {
                if (c > 0 && c < cols - 1) && 
                abs(board[r][c]) == abs(board[r][c+1]) && 
                abs(board[r][c]) == abs(board[r][c-1]) { 
                    negVal := -(abs(board[r][c]))
                    board[r][c], board[r][c+1], board[r][c-1] = negVal, negVal, negVal
                    streakExists = true
                }
                if (r > 0 && r < rows - 1) &&
                abs(board[r][c]) == abs(board[r+1][c]) && 
                abs(board[r][c]) == abs(board[r-1][c]) {
                    negVal := -(abs(board[r][c]))
                    board[r][c], board[r+1][c], board[r-1][c] = negVal, negVal, negVal
                    streakExists = true
                }
            }
        }
    }
    return streakExists
}

func crushAndDrop(board [][]int) {
    rows := len(board)
    cols := len(board[0])

    for c := range cols {
        shiftedList := make([]int, rows)
        index := 0
        for r := rows - 1; r >= 0 && board[r][c] != 0; r-- {
            if board[r][c] > 0 {
                shiftedList[index] = board[r][c]
                index++
            }
        }
        index = 0
        for r := rows - 1; r >= 0 ; r-- {
            board[r][c] = shiftedList[index]
            index++
        }
    }
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
