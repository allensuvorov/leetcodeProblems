func countGoodNodes(edges [][]int) int {
    goodNodesCount := 0
    children := make([][]int, len(edges) + 1)
    for _, edge := range edges {
        children[edge[0]] = append(children[edge[0]], edge[1])
        children[edge[1]] = append(children[edge[1]], edge[0])
    }

    visited := map[int]bool{0:true}
    var dfs func(int) int
    dfs = func(node int) int {
        prevSize := 0
        isGood := true
        sum := 0
        for _, child := range children[node] { 
            if !visited[child] {
                visited[child] = true
                size := dfs(child)
                if prevSize != 0 && size != prevSize {
                    isGood = false
                }
                sum += size
                prevSize = size
            }
        }
        if isGood {
            goodNodesCount++
        }
        return sum + 1
    } 

    dfs(0)
    return goodNodesCount
}
