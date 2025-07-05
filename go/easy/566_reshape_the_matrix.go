func matrixReshape(mat [][]int, r int, c int) [][]int {
    if r * c != len(mat) * len(mat[0]) {
        return mat
    }

    res := make([][]int, r)
    for row := range r {
        res[row] = make([]int, c)
    }

    for row := range r {
        for col := range c {
            i := row * c + col
            res[row][col] = mat[i / len(mat[0])][i % len(mat[0])]
        }
    }
    return res
}
