func countGoodNodes(edges [][]int) int {
    goodNodesCount := 0
    children := make(map[int][]int)
    for _, edge := range edges {
        parent := edge[0]
        child := edge[1]
        for range 2 {
            if _, ok := children[parent]; !ok {
                children[parent] = []int{}
            }
            children[parent] = append(children[parent], child)
            parent, child = child, parent
        }
    }

    visited := map[int]bool{0:true}
    var dfs func(int) int
    dfs = func(node int) int {
        size := 0
        prevSize := 0
        isGood := true
        sum := 0
        for _, child := range children[node] { 
            if !visited[child] {
                visited[child] = true
                size = dfs(child)
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
