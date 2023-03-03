func checkValid(matrix [][]int) bool {
    
    nums := make([]int, len(matrix))

    for i := range matrix {
        for j := range matrix{
            nums[matrix[i][j]-1]++
            if nums[matrix[i][j]-1] != i + 1 {
                return false
            }
        }
    }
    
    for j := range matrix {
        for i := range matrix{
            nums[matrix[i][j]-1]++
            if nums[matrix[i][j]-1] != len(matrix) + j + 1 {
                return false
            }
        }
    }
    return true
}
