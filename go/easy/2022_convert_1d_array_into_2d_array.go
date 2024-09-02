func construct2DArray(original []int, m int, n int) [][]int {
    if len(original) != m * n {
        return [][]int{}
    }
    res := make([][]int, m)
    
    i := 0
    for r := range res {
        row := make([]int, n)
        for c := range row {
            row[c] = original[i]
            i++
        }
        res[r] = row
    }
    return res
}
