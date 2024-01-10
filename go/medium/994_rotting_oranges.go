func orangesRotting(grid [][]int) int {
    // get initial 2s
    minutes := 0
    q := [][]int{}
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 2 {
                q = append(q, []int{i,j,0})
            }
        }
    }
    // BFS
    for len(q) > 0 {
        for _, edge := range [][]int{[]int{1,0}, []int{-1,0}, []int{0,1}, []int{0,-1}} {
            nr := q[0][0] + edge[0]
            nc := q[0][1] + edge[1]
            if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nr][nc] == 1 {
                grid[nr][nc] = 2
                minutes = q[0][2] + 1
                q = append(q, []int{nr, nc, minutes})
            }
        }
        q = q[1:]
    }
    // look for 1s
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                fmt.Println("fresh oranges left")
                return -1
            }
        }
    }

    return minutes
}
