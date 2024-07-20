func allPathsSourceTarget(graph [][]int) [][]int {
    destination := len(graph) - 1
    origin := 0
    ans := [][]int{}

    var dfs func(path []int)
    dfs = func(path []int) {
        now := path[len(path) - 1]
        if now == destination {
            temp := make([]int, len(path))
            copy(temp, path)
            ans = append(ans, temp)
        } else {
            for _, child := range graph[now] {
                dfs(append(path, child))
            }
        }
    }
    
    dfs([]int{origin})
    return ans
}
