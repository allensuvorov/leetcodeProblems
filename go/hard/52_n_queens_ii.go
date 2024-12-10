func totalNQueens(n int) int {
    solutionCount := 0
    var backTrack func (chessboard [][]int)
    backTrack = func (chessboard [][]int) {
        if len(chessboard) == 0 {
            solutionCount++
            return
        }

        for c := range n {
            if chessboard[0][c] == 0 {
                clone := cloneChessboard(chessboard)
                placeQueen(clone, c)
                backTrack(clone[1:]) // 3
            }
        }
    }
    
    backTrack(initChessboard(n, n))
    return solutionCount
}

func cloneChessboard(chessboard [][]int) [][]int {
    clone := initChessboard(len(chessboard), len(chessboard[0]))
    for r := range chessboard {
        for c := range chessboard[0] {
            clone[r][c]= chessboard[r][c]
        }
    }
    return clone
}

func initChessboard (r, c int) [][]int{
    chessboard := make([][]int, r)
    for r := range r {
        chessboard[r] = make([]int, c)
    }
    return chessboard
}

func placeQueen (chessboard [][]int, col int) {
    rows, cols := len(chessboard), len(chessboard[0])
    // queen or attack - 1
    // attack down
    attack := col + 1
    for r, c := 0, col; r < rows; r++  {
        chessboard[r][c] = attack
    }
    // attack \
    for r, c := 0, col; r < rows && c < cols; {
        chessboard[r][c] = attack
        r++
        c++
    }
    // attack /
    for r, c := 0, col; r < rows && c >= 0; {
        chessboard[r][c] = attack
        r++
        c--
    }
}

