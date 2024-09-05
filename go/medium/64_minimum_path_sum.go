func minPathSum(grid [][]int) int {
    dp := make(map[[2]int]int)

    rows := len(grid)
    cols := len(grid[0])

    q := [][2]int{[2]int{0,0}}
    for len(q) > 0 {
        diagonalSize := len(q)
        for range diagonalSize {
            now := q[0]
            q = q[1:]

            up := -1
            if now[0] - 1 >= 0 {
                parent := [2]int{now[0] - 1, now[1]}
                up = dp[parent]
            }

            left := -1
            if now[1] - 1 >= 0 {
                parent := [2]int{now[0], now[1] - 1}
                left = dp[parent]
            }

            switch {
            case up > 0 && left > 0:
                dp[now] = min(up, left)
            case up > 0:
                dp[now] = up
            case left > 0:
                dp[now] = left
            }

            dp[now] += grid[now[0]][now[1]]

            if now[0] + 1 < rows {
                if len(q) == 0 || q[len(q)-1] != [2]int{now[0] + 1, now[1]} {
                    q = append(q, [2]int{now[0] + 1, now[1]})
                }
            }
            if now[1] + 1 < cols {
                q = append(q, [2]int{now[0], now[1] + 1})
            }
        }
    }
    end := [2]int{rows-1, cols-1}
    return dp[end]
}
