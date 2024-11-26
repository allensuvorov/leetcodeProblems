func snakesAndLadders(board [][]int) int {
    n := len(board)
    moves := 0
    q := []int{1}

    for len(q) > 0 {
        size := len(q)
        for range size {
            now := q[0]
            q = q[1:]
            if now == n*n {
                return moves
            }
            for next := now + 1; next <= n*n && next <= now + 6; next++ {
                r, c := getCoordinates(next, n)
                switch board[r][c] {
                case -1: 
                    q = append(q, next)
                case -2:
                    continue
                default:
                    q = append(q, board[r][c])
                }
                board[r][c] = -2 // mark visited with -2
            }
        }
        moves++
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
