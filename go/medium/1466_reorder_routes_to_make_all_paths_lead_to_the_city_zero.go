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
