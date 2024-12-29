func equalPairs(grid [][]int) int {
    n := len(grid)
    const maxLen = 200
    // count rows and cols
    rowHashCount := make(map[[maxLen]int]int)
    for r := range n {
        arr := [maxLen]int{}
        for c := range n {
            arr[c] = grid[r][c]
        }
        rowHashCount[arr]++
    }

    pairsCount := 0
    
    for c := range n {
        arr := [maxLen]int{}
        for r := range n {
            arr[r] = grid[r][c]
        }
        pairsCount += rowHashCount[arr]
    }

    return pairsCount
}
