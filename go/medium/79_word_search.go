func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
    
    var dfs func(int, int, int) bool
    dfs = func(r, c, i int) bool {
        if i == len(word) {
            return true
        }
        if r < 0 || r == rows || c < 0 || c == cols || board[r][c] != word[i]{
            return false
        }

        tmp := board[r][c]
        board[r][c] = '1'
        res := 
            dfs(r + 1, c, i + 1) ||
            dfs(r - 1, c, i + 1) ||
            dfs(r, c + 1, i + 1) ||
            dfs(r, c - 1, i + 1)
        board[r][c] = tmp
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
