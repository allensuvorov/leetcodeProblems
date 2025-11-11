func equalPairs(grid [][]int) int {
    n := len(grid)
    const keySize = 200
    rowCount := map[[keySize]int]int{}
    for r := range n {
        var temp [keySize]int
        for c := range n {
            temp[c] = grid[r][c]
        }
        rowCount[temp]++
    }

    res := 0
    for c := range n {
        var temp [keySize]int
        fmt.Println(temp)
        for r := range n {
            temp[r] = grid[r][c]
        }
        res = res + rowCount[temp]
    }
    return res
}
