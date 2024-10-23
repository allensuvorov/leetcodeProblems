func equalPairs(grid [][]int) int {
    // O(n^2) or O(n^3) time complexity
    n := len(grid)
    const maxLen = 200
    rowInventory := make(map[[maxLen]int]int, n)

    for r := range n {
        row := [maxLen]int{}
        for c := range n {
            row[c] = grid[r][c]
        }
        rowInventory[row]++
    }

    equalPairsCount := 0
    
    for c := range n {
        col := [maxLen]int{}
        for r := range n {
            col[r] = grid[r][c]
        }
        equalPairsCount += rowInventory[col]
    }

    return equalPairsCount
}
