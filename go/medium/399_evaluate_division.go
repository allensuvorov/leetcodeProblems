func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    adj := make(map[string][]int)
    for i, v := range equations {
        adj[v[0]] = append(adj[v[0]], i)
        adj[v[1]] = append(adj[v[1]], i)
    }

    visited := make([]bool, len(equations))
    var dfs func(a, b string) float64
    dfs = func(a, b string) float64 {
        if a == b {
            return 1.0
        }
        for _, i := range adj[a] {
            if !visited[i] {
                visited[i] = true
                res := 1.0
                if a == equations[i][0] {
                    res = dfs(equations[i][1], b) * values[i]
                } else {
                    res = dfs(equations[i][0], b) / values[i]
                }
                if res > 0 {
                    return res
                }
            }
        } 
        return -1.0
    }

    res := make([]float64, len(queries))
    for i, v := range queries {
        if adj[v[0]] == nil || adj[v[1]] == nil{
            res[i] = -1.0
        } else {
            visited = make([]bool, len(equations))
            res[i] = dfs(v[0], v[1])
        }
    }
    return res
}
