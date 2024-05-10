func matrixSum(nums [][]int) int {
    for i := range nums {
        sort.Ints(nums[i])
    }
    score := 0
    for j := len(nums[0]) - 1; j >= 0; j -- {
        maxCol := 0
        for i := range nums {
            maxCol = max(nums[i][j], maxCol)
        }
        score += maxCol
    }
    return score
}
