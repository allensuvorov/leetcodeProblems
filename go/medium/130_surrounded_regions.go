func solve(board [][]byte)  {
    rows, cols := len(board), len(board[0])
    
    // DFS to ID border regions, mark them with 'T'
    var dfs func(r, c int) 
    
    dfs = func(r, c int) {
        if r < 0 || c < 0 || r == rows || c == cols || board[r][c] != 'O' {
            return
        }
        board[r][c] = 'T'
        dfs(r-1, c)
        dfs(r+1, c)
        dfs(r, c+1)
        dfs(r, c-1)
    }
    
    // loop over border
    for c := 0; c < cols; c++ {
        dfs(0, c)
    }
    for c := 0; c < cols; c++ {
        dfs(rows-1, c)
    }
    for r := 1; r < rows - 1; r++ {
        dfs(r, 0)
    }
    for r := 1; r < rows - 1; r++ {
        dfs(r, cols-1)
    }

    // loop over the matrix, flip 'O' to 'X', and 'T' to 'O'
    for r := range rows {
        for c := range cols {
            if board[r][c] == 'O' {
                board[r][c] = 'X'
            }
            if board[r][c] == 'T' {
                board[r][c] = 'O'
            }
        }
    }
}
