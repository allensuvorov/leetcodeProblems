func rowAndMaximumOnes(mat [][]int) []int {
    maxCnt := 0
    maxPos := 0

    for i := range mat{
        count := 0
        for j := range mat[i]{
            count += mat[i][j]
        }
        if maxCnt < count {
            maxCnt = count
            maxPos = i
        }
    }
    return []int{maxPos, maxCnt}
}
