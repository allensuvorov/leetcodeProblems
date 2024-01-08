func nearestExit(maze [][]byte, entrance []int) int {  
    maze[entrance[0]][entrance[1]] = '1'
    q := [][]int{append(entrance, 0)}
    for len(q) > 0 {
        for _, edge := range [][]int{[]int{-1,0}, []int{1,0}, []int{0,1}, []int{0,-1}} {
            nextRow := q[0][0] + edge[0]
            nextCol := q[0][1] + edge[1]
            if nextRow >= 0 && nextRow < len(maze) && nextCol >= 0 && nextCol < len(maze[0]) && maze[nextRow][nextCol] == '.' {
                if nextRow == 0 || nextRow == len(maze) - 1 || nextCol == 0 || nextCol == len(maze[0]) - 1  {
                    return q[0][2] + 1
                }
                maze[nextRow][nextCol] = '1'
                q = append(q, []int{nextRow, nextCol, q[0][2] + 1})
            }
        }
        q = q[1:]
    }
    return -1
}
