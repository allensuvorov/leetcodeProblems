//118. Pascal's Triangle
func generate(numRows int) [][]int {
    var triangle = make([][]int, numRows)
    for row:= 0; row < numRows; row++{
        triangle[row] = make([]int, row+1)
        
        triangle[row][0] = 1
        triangle[row][row] = 1        
        
        for j:=1; j < row; j++{
            triangle[row][j] = triangle[row-1][j-1] + triangle[row-1][j]
        }
    }
    return triangle
}
