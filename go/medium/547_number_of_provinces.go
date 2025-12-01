// recurcive BFS
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

// iterative DFS
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    result := 0
    seen := make(map[int]bool)
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
            for nei := range isConnected[now] {
                if !seen[nei] && isConnected[now][nei] == 1 {
                    todo = append(todo, nei)
                }
            }
        }
    }
    return result
}
