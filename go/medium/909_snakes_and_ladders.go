func snakesAndLadders(board [][]int) int {
    n := len(board)
    q := []int{1}

    for moves := 0; len(q) > 0; moves++ {
        size := len(q)
        for range size {
            now := q[0]
            q = q[1:]
            if now == n*n {
                return moves
            }
            for next := now + 1; next <= min(n*n, now + 6); next++ {
                r, c := getCoordinates(next, n)
                if board[r][c] != -2 {
                    if board[r][c] == -1 {
                        q = append(q, next)
                    } else {
                        q = append(q, board[r][c])
                    }
                    board[r][c] = -2 // mark visited with -2
                }
            }
        }
    }
    return -1
}

func getCoordinates(label, n int) (int, int) {
    r := n - 1 - (label - 1) / n
    c := (label - 1) % n
    if ((label - 1) / n) % 2 == 1 {
        c = n - 1 - c
    }
    return r, c
}
