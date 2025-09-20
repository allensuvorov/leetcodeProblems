func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    res := make([]int, len(queries))
    graph := make([][]int, n) // adjacency table
    // initialize the adj table
    for node := 0; node < len(graph) - 1; node++ {
        graph[node] = append(graph[node], node + 1)
    }
    for i, v := range queries {
        graph[v[0]] = append(graph[v[0]], v[1]) // add new route
        res[i] = BFS(graph, n)
    }
    return res
}

func BFS(graph [][]int, n int) int {
    q := []int{0} // queue
    seen := map[int]bool{0:true}
    dist := 0

    for len(q) > 0 {
        size := len(q)
        for range size {
            now := q[0]
            q = q[1:]
            seen[now] = true
            if now == n - 1 {
                return dist
            }

            for _, nei := range graph[now] {
                if !seen[nei] {
                    q = append(q, nei)
                }
            }
        }
        dist++
    }
    return dist
}
