func solve(board [][]byte)  {
    rows, cols := len(board), len(board[0])
    // loop over the matrix, excluding outer edges
    for r := 1; r < rows - 1; r++ {
        for c := 1; c < cols - 1; c++ {
            if board[r][c] == 'O' {
                visited := make(map[[2]int]bool)
                isSurrounded, visited := dfs(r, c, board, visited)
                if isSurrounded {
                    capture(board, visited)
                }
            }
        }
    }
}

// DFS over each region, track visited
func dfs(r, c int, board [][]byte, visited map[[2]int]bool) (bool, map[[2]int]bool) {
    if board[r][c] == 'X' || visited[[2]int{r,c}] {
        return true, visited
    }

    if r == 0 || c == 0 || r == len(board) - 1 || c == len(board[0]) -1 {
        return false, visited
    }

    visited[[2]int{r, c}] = true

    if isSurrounded, visited := dfs(r-1, c, board, visited); !isSurrounded {
        return false, visited
    }
    if isSurrounded, visited := dfs(r+1, c, board, visited); !isSurrounded {
        return false, visited
    }
    if isSurrounded, visited := dfs(r, c-1, board, visited); !isSurrounded {
        return false, visited
    }
    if isSurrounded, visited := dfs(r, c+1, board, visited); !isSurrounded {
        return false, visited
    }
    return true, visited
}

// if region surrounded, capture visited
func capture(board [][]byte, visited map[[2]int]bool) {
    for pos := range visited {
        board[pos[0]][pos[1]] = 'X'
    }
}
