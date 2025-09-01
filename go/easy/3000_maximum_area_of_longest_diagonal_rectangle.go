func areaOfMaxDiagonal(dimensions [][]int) int {
    var maxDiagonalRectangleArea int
    var maxDiagonal float64
    for _, v := range dimensions {
        diagonal := math.Sqrt(float64(v[0]*v[0] + v[1]*v[1]))
        if diagonal > maxDiagonal {
            maxDiagonal = diagonal
            maxDiagonalRectangleArea = v[0] * v[1]
        } else if diagonal == maxDiagonal {
            maxDiagonalRectangleArea = max(maxDiagonalRectangleArea, v[0] * v[1])
        }
    }
    return maxDiagonalRectangleArea
}
