func numIslands(grid [][]byte) int {
    // look for and island - find 1s and count it
    // traverse the island and mark it - change 1 to 2
    count := 0
    var dfs func(i, j int) 
    dfs = func(i, j int){
        if grid[i][j] == '1' {
            grid[i][j] = '2' 
            if i > 0 { dfs(i-1, j) }
            if i < len(grid) - 1 { dfs(i+1,j) }
            if j > 0 { dfs(i, j-1) }
            if j < len(grid[0])-1 {dfs(i, j+1) }
        }
    }
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == '1' {
                count++
                dfs(i, j)
            }
        }
    }
    return count
}
