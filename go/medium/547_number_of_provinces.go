func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    provCount := 0
    for i := range n {
        if isConnected[i][i] == 1 {
            provCount++
            bfs(isConnected, i)
        }
    }
    return provCount
}

func bfs(isConnected [][]int, now int) {
    q := []int{now}
    for len(q) > 0 {
        now := q[0]
        q = q[1:]
        isConnected[now][now] = 2
        for nei := range isConnected {
            if isConnected[now][nei] == 1 && isConnected[nei][nei] == 1 {
                q = append(q, nei)
            }
        }
    }
}

// build graph, iterative dfs
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    // DFS, track seen
    // every !seen is new province
    result := 0
    seen := make(map[int]bool)

    // build adjList
    g := make([][]int, n)
    for i := range n {
        for j := range n {
            if i != j && isConnected[i][j] == 1 {
                g[i] = append(g[i], j)
            }
        }
    }
    
    for city := range n {
        if !seen[city] {
            result++
        }
        // dfs
        todo := []int{city}
        for len(todo) > 0 {
            now := todo[len(todo)-1]
            todo = todo[:len(todo)-1]
            seen[now] = true
            for _, nei := range g[now] {
                if !seen[nei] {
                    todo = append(todo, nei)
                }
            }
        }
    }
    return result
}
