func minReorder(n int, connections [][]int) int {
    origins := make(map[int]map[int]int, n)
    for i, v := range connections {
        if _, ok := origins[v[0]]; !ok {
            origins[v[0]] = make(map[int]int)
        }
        if _, ok := origins[v[1]]; !ok {
            origins[v[1]] = make(map[int]int)
        }
        origins[v[0]][v[1]] = i
        origins[v[1]][v[0]] = i
    }
    visited := make(map[int]bool, n)
    reorderCount := 0

    var dfs func(destination int)
    dfs = func(destination int) {
        if visited[destination] {
            return
        }
        visited[destination] = true
        for origin, i := range origins[destination] {
            if !visited[origin] && origin != connections[i][0] {
                reorderCount++
            }
            dfs(origin)
        }
    }
    dfs(0)
    return reorderCount
}
