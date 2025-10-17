func equalPairs(grid [][]int) int {
    const n = 200
    rowCounts := make(map[[n]int]int) 
    res := 0

    for r := range grid {
        rowArray := [n]int{}
        for i, v := range grid[r] {
            rowArray[i] = v
        }
        rowCounts[rowArray]++
    }
    
    for c := range grid[0] {
        colArray := [n]int{}
        for r := range grid[c] {
            colArray[r] = grid[r][c]
        }
        res = res + rowCounts[colArray]
    }
    return res
}
