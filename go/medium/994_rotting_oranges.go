func orangesRotting(grid [][]int) int {
    minutes := 0
    fresh := 0
    q := [][]int{}
    for i := range grid {
        for j := range grid[i] {
            switch grid[i][j] {
            case 1:
                fresh++
            case 2:
                q = append(q, []int{i,j,0})
            }
        }
    }
    if fresh == 0 {
        return 0
    }

    // BFS
    for len(q) > 0 {
        for _, edge := range [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
            nr := q[0][0] + edge[0]
            nc := q[0][1] + edge[1]
            if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nr][nc] == 1 {
                grid[nr][nc] = 2
                fresh--
                minutes = q[0][2] + 1
                q = append(q, []int{nr, nc, minutes})
            }
        }
        q = q[1:]
    }
    if fresh != 0 {
        return -1
    }
    return minutes
}
