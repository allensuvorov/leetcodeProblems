func matrixReshape(mat [][]int, r int, c int) [][]int {
    rows := len(mat)
    cols := len(mat[0])

    if r * c != rows * cols {
        return mat
    }

    res := make([][]int, r)
    for row := range res {
        res[row] = make([]int, c)
    }

    for targetR := range r {
        for targetC := range c {
            i := targetR * c + targetC
            res[targetR][targetC] = mat[i / cols][i % cols]
        }
    }
    return res
}
