func diagonalSum(mat [][]int) int {
    sum := 0
    l, r := 0, len(mat) - 1
    for i := range mat {
        sum += mat[i][l]
        if i != r {
            sum += mat[i][r]
        }
        l++
        r--
    }
    return sum
}
