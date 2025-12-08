// v1
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

// v2
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    varSet := map[string]bool{}
    for _, v := range equations {
        varSet[v[0]] = true
        varSet[v[1]] = true
    }
    var getMultiplier func(evaluated, unit string, seen map[string]bool) float64
    getMultiplier = func(evaluated, unit string, seen map[string]bool) float64 {
        if evaluated == unit {
            return 1
        }
        seen[evaluated] = true // a
        nextEvaluated := ""

        for i, e := range equations { // check all neighbours
            multiplier := -1.0
            if e[0] == evaluated && !seen[e[1]]{
                multiplier = values[i]
                nextEvaluated = e[1]
            } else if e[1] == evaluated && !seen[e[0]]{
                multiplier = 1 / values[i]
                nextEvaluated = e[0]
            }
            
            if multiplier != -1.0 {
                nextMultiplier := getMultiplier(nextEvaluated, unit, seen)
                if nextMultiplier == -1.0 {
                    continue
                } else {
                    return multiplier * nextMultiplier
                }
            }
        }
        return -1.0
    }
    
    res := make([]float64, len(queries))
    for i, v := range queries {
        if varSet[v[0]] && varSet[v[1]] {
            evaluated := v[0]
            unit := v[1] // common denominator
            seen := map[string]bool{}
            res[i] = getMultiplier(evaluated, unit, seen)
        } else {
            res[i] = -1
        }
    }
    return res
}
