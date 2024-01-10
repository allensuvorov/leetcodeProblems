func nearestExit(maze [][]byte, entrance []int) int {  
    maze[entrance[0]][entrance[1]] = '1'
    q := [][]int{append(entrance, 0)}
    for len(q) > 0 {
        for _, edge := range [][]int{{-1,0}, {1,0}, {0,1}, {0,-1}} {
            nr := q[0][0] + edge[0]
            nc := q[0][1] + edge[1]
            if nr >= 0 && nr < len(maze) && 
                nc >= 0 && nc < len(maze[0]) && 
                maze[nr][nc] == '.' {
                if nr == 0 || nr == len(maze) - 1 || 
                    nc == 0 || nc == len(maze[0]) - 1  {
                    return q[0][2] + 1
                }
                maze[nr][nc] = '1'
                q = append(q, []int{nr, nc, q[0][2] + 1})
            }
        }
        q = q[1:]
    }
    return -1
}
