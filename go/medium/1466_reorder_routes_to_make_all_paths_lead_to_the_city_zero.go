func minReorder(n int, connections [][]int) int {
    adj := make([][][]int, n)
    for _, v := range connections {
        adj[v[0]] = append(adj[v[0]], []int{v[1], 1}) // 0 - from
        adj[v[1]] = append(adj[v[1]], []int{v[0], 0}) // 1 - to
    }

    count := 0
    var dfs func(node, parent int)
    dfs = func(node, parent int) {
        for _, v := range adj[node] {
            nei, direction := v[0], v[1]
            if nei != parent {
                count += direction
                dfs(nei, node)
            }
        }
    }
    dfs(0, -1)
    return count
}

// explicit
func minReorder(n int, connections [][]int) int {
    // build undirected graph adjList with a direction parameter (1 - reversed, 0 - not reversed)
    // DFS and count sum of those direction parameters.
    g := make([][]pair, n) // city -> list of pairs
    for _, v := range connections {
        a := v[0]
        b := v[1]
        g[a] = append(g[a], pair{b, 1})
        g[b] = append(g[b], pair{a, 0})
    }
    seen := make([]bool, n)
    res := 0
    
    for stack := []int{0}; len(stack) > 0;  {
        now := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        seen[now] = true
        for _, nei := range g[now] {
            if !seen[nei.city] {
                stack = append(stack, nei.city)
                res += nei.direction
            }
        }
    }
    return res
}

type pair struct {
    city int
    direction int // 1 - reversed, 0 - not reversed
}
