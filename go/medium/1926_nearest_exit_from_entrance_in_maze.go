func nearestExit(maze [][]byte, entrance []int) int {
    dist := 0
    entrance = append(entrance, dist)
    q := [][]int{entrance}
    maze[entrance[0]][entrance[1]] = '1'
    
    for len(q) != 0 {
        now := q[0]
        q = q[1:]
        dist = now[2]
        
        checkVertex := func(m, n int) bool {
            if m >= 0 && m < len(maze) && n >= 0 && n < len(maze[0]) {
                // empty - mark dist and enq
                if maze[m][n] == '.' {
                    q = append(q, []int{m, n, dist + 1})
                    maze[m][n] = '1'
                }
            } else if dist != 0 { // now is an exit
                return true
            }
            return false
        }

        checkNearestVertices := func(now []int) bool {
            if exit := checkVertex(now[0]-1, now[1]); exit { return true } // up 
            if exit := checkVertex(now[0]+1, now[1]); exit { return true } // down
            if exit := checkVertex(now[0], now[1]+1); exit { return true } // right
            if exit := checkVertex(now[0], now[1]-1); exit { return true } // left
            return false
        }
        
        if exit := checkNearestVertices(now); exit {
            return int(dist)
        }
    }
    
    return -1
}
