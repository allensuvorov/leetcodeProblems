func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
    path := make([][]bool, rows)
    for i := range path {
        path[i] = make([]bool, cols)
    }
    
    var dfs func(r, c, i int) bool
    dfs = func(r, c, i int) bool {
        if i == len(word) {
            return true
        }
        if r < 0 || r == rows || c < 0 || c == cols || path[r][c] || board[r][c] != word[i] {
            return false
        }
        path[r][c] = true
        res := 
            dfs(r + 1, c, i + 1) ||
            dfs(r - 1, c, i + 1) ||
            dfs(r, c + 1, i + 1) ||
            dfs(r, c - 1, i + 1)
        path[r][c] = false
        return res
    }

    for r := range board {
        for c := range board[r] {
            if dfs(r, c, 0) {
                return true
            }
        }
    }
    return false
}
