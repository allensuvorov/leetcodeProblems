func queensAttacktheKing(queens [][]int, king []int) [][]int {
    var (
        queenSet map[[2]int]bool = make(map[[2]int]bool)
        res [][]int
    )

    for _, v := range queens {
        queenSet[[2]int{v[0],v[1]}] = true
    }

    directions := [][]int{{1,0},{-1,0},{0,1},{0,-1},{1,1},{-1,-1},{1,-1},{-1,1}}

    for _, dir := range directions {
        for r, c := king[0], king[1]; r >= 0 && r < 8 && c >= 0 && c < 8; {
            if queenSet[[2]int{r,c}] {
                res = append(res, []int{r,c})
                break
            }
            r += dir[0]
            c += dir[1]
        }
    }
    return res
}
