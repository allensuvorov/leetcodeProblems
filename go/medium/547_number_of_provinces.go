func findCircleNum(isConnected [][]int) int {
    count := 0
    visited := make([]bool, len(isConnected))
    for i := range isConnected {
        if !visited[i] {
            count++
            dfs(isConnected, visited, i)
        }
    }
    return count
}

func dfs(isConnected [][]int, visited []bool, now int) {
    if !visited[now] {
        visited[now] = true
        for i := range isConnected {
            if isConnected[now][i] == 1 && !visited[i] {
                dfs(isConnected, visited, i)
            }
        }
    }
}
