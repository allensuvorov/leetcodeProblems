func generate(numRows int) [][]int {
    var triangle [][]int
    for row:= 0; row < numRows; row++{
        triangle = append([row][0]int{1},triangle...)
        
        // triangle[row][0] = 1
        // triangle[row][row] = 1        
        
        for j:=1; j < row; j++{
            triangle[row][j] = triangle[row-1][j-1] + triangle[row-2][j]
        }
        triangle = append(triangle, [row][row]int{1})
    }
    return triangle
}
