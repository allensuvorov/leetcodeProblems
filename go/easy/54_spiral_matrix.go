func spiralOrder(matrix [][]int) []int {
    const visited int = 101
    rows, cols := len(matrix), len(matrix[0])
    size := rows*cols
    ans := make([]int, 0, rows * cols)
    dirs := [][]int{{0,1}, {1,0}, {0,-1}, {-1,0}}
    nowR, nowC := 0, 0
    ans = append(ans, matrix[nowR][nowC])
    matrix[nowR][nowC] = visited
    dir := dirs[0]
    for len(ans) < size { 
        nextR, nextC := nowR + dir[0], nowC + dir[1]; 
        for nextR >= 0 && nextR < rows && nextC >= 0 && nextC < cols && matrix[nextR][nextC] != visited {
            ans = append(ans, matrix[nextR][nextC])
            nowR, nowC = nextR, nextC
            matrix[nowR][nowC] = visited
            nextR += dir[0]
            nextC += dir[1]
        }
        dir = getNextDir(dir)
    }
    return ans
}

func getNextDir (curDir []int) []int {
    switch {
    case curDir[0] == 0 && curDir[1] == 1: return []int{1,0}
    case curDir[0] == 1 && curDir[1] == 0: return []int{0,-1}
    case curDir[0] == 0 && curDir[1] == -1: return []int{-1,0}
    case curDir[0] == -1 && curDir[1] == 0: return []int{0,1}
    }
    return nil
}
