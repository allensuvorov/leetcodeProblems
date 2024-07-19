func luckyNumbers (matrix [][]int) []int {
    mins := map[int]bool{}
    ans := []int{}
    for i := range matrix{
        curMin := matrix[i][0]
        for j := range matrix[i] {
            curMin = min(curMin, matrix[i][j])
        }
        mins[curMin] = true
    }
    for j := range matrix[0] {
        curMax := matrix[0][j]
        for i := range matrix {
            curMax = max(curMax, matrix[i][j])
        }
        if mins[curMax] {
            ans = append(ans, curMax)
        }
    }
    return ans
}
