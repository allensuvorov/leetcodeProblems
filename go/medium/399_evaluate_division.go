func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    adj := map[string][]int{}

    for i := range equations {
        adj[equations[i][0]] = append(adj[equations[i][0]],i)
        adj[equations[i][1]] = append(adj[equations[i][1]],i)
    }

    res := make([]float64, len(queries))
    visited := make([]bool, len(equations))

    var dfs func(a, b string) float64
    dfs = func(a, b string) float64{
        if a == b {
            return 1.0
        }

        for _, v := range adj[a] {
            if !visited[v] {
                visited[v] = true
                if a == equations[v][0] {
                    res := dfs(equations[v][1], b) * values[v]
                    if res > 0 {
                        return res
                    }
                } else {
                    res := dfs(equations[v][0], b) / values[v]
                    if res > 0 {
                        return res
                    }
                }
            }
        }
        return -1.0
    }

    for i, v:= range queries{
        if adj[v[0]] == nil || adj[v[1]] == nil {
            res[i] = -1.0
        } else {
            visited = make([]bool, len(equations))
            res[i] = dfs(v[0], v[1])
        }
    }
    return res
}