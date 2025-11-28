func earliestAcq(logs [][]int, n int) int {
    // binary search    bbbbgggg
    slices.SortFunc(logs, func(a, b []int) int{
        return a[0] - b[0]
    })

    l, r := 0, len(logs)
    for l < r {
        m := l + (r-l)/2
        // build graf from 1 to middle
        g := map[int][]int{}
        buildGraph(g, logs[:m+1])
        seen := map[int]bool{}
        p := logs[0][1]
        dfs(p, g, seen) // linear+

        if len(seen) == n { // all n are friends
            r = m
        } else {
            l = m + 1
        }
    }
    if l == len(logs) {
        return -1
    }
    return logs[l][0]
}

func buildGraph(g map[int][]int, logs [][]int) {
    for i := range logs {
        p1 := logs[i][1]
        p2 := logs[i][2]
        g[p1] = append(g[p1], p2)
        g[p2] = append(g[p2], p1)
    }
}

// check the graph has all people
func dfs(p int, g map[int][]int, seen map[int]bool) {
    seen[p] = true
    for _, nei := range g[p] {
        if !seen[nei] {
            dfs(nei, g, seen)
        }
    }
}
