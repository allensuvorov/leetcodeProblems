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
                if maze[m][n] == '.' {  // empty - mark dist and enq
                    q = append(q, []int{m, n, dist + 1})
                    maze[m][n] = '1'
                }
            } else if dist != 0 { // now is an exit
                return true
            }
            return false
        }

        checkNearestVertices := func(now []int) bool {
            edges := [][]int{[]int{-1,0}, []int{1,0}, []int{0,1}, []int{0,-1}}   
            for _, v := range edges{
                if exit := checkVertex(now[0] + v[0], now[1] + v[1]); exit { 
                    return true 
                }
            }
            return false
        }
        
        if exit := checkNearestVertices(now); exit {
            return int(dist)
        }
    }
    
    return -1
}
