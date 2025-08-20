func nearestExit(maze [][]byte, entrance []int) int {
    q := [][]int{entrance}
    stepCount := 0
    rows := len(maze)
    cols := len(maze[0])
    // mark visited 
    maze[entrance[0]][entrance[1]] = '-'
    directions := [][]int{{-1,0}, {1,0}, {0,1}, {0,-1}}

    for len(q) > 0 {
        zoneSize := len(q)
        for range zoneSize {
            now := q[0]
            q = q[1:]
            
            // is this exit?
            if stepCount > 0 && (now[0] == 0 || now[1] == 0 || now[0] == rows - 1 || now[1] == cols - 1) {
                return stepCount
            }
            // add next steps to the q
            for _, direction := range directions {
                nextRow := now[0] + direction[0]
                nextCol := now[1] + direction[1]
                if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols {
                    if maze[nextRow][nextCol] == '.' {
                        // mark as enqued already
                        maze[nextRow][nextCol] = '-'
                        q = append(q, []int{nextRow, nextCol})
                    }
                }
            }
        }
        stepCount++
    }
    return -1
}

// BFS with a q
