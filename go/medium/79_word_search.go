func exist(board [][]byte, word string) bool {
    for i := range board {
        for j := range board[i] {
            if board[i][j] == word[0] {
                visited := make([][]bool, len(board))
                for i := range visited {
                    visited[i] = make([]bool, len(board[0]))
                }

                if dfs(board, word, visited, []int{i, j}, 1) {
                    return true
                }
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, visited [][]bool, now []int, i int) bool {
    if i == len(word) {
        return true
    }
    visited[now[0]][now[1]] = true

    for _, direction := range [][]int{{1,0}, {-1,0}, {0, 1}, {0, -1}} {
        nr := now[0] + direction[0]
        nc := now[1] + direction[1]
        if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) && 
            !visited[nr][nc] &&
            board[nr][nc] == word[i] {
                if dfs(board, word, visited, []int{nr, nc}, i + 1) {
                    return true
                }
        }
    }
    visited[now[0]][now[1]] = false
    return false
}
