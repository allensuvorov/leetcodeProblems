func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    for r := n - 2; r >= 0; r-- {
        for i := range triangle[r] {
            triangle[r][i] += min(
                triangle[r+1][i], 
                triangle[r+1][i+1],
            ) 
        }
    }
    return triangle[0][0]
}
