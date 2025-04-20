/*
* tc Big O(n)
* sc Big O(n)
* patterns: prefixSum, most recent position, dfs
*/
func longestSpecialPath(edges [][]int, nums []int) []int {
    n := len(nums)
    
    g := make([][][]int, n) // build graph (adjacency list)
    for _, edge := range edges {
        a, b, length := edge[0], edge[1], edge[2]
        g[a] = append(g[a], []int{b, length})
        g[b] = append(g[b], []int{a, length})
    }

    last := make(map[int]int) // last position of a value
    for _, v := range nums {
        last[v] = -1
    }

    prefSum := make([]int, n)
    var maxLen, minNodeCount int
    path := []int{} // nodes in a path

    var dfs func(node, tail, sum int)
    dfs = func(node, tail, sum int) {
        prefSum[node] = sum
        
        if last[nums[node]] >= tail { // move the tail 
            tail = last[nums[node]] + 1
        }
        old := last[nums[node]]
        last[nums[node]] = len(path)
        path = append(path, node)

        currLen := sum - prefSum[path[tail]]
        if currLen > maxLen {
            maxLen = currLen
            minNodeCount = len(path) - tail
        } else if currLen == maxLen {
            minNodeCount = min(minNodeCount, len(path) - tail)
        }

        for _, nei := range g[node] {
            next, length := nei[0], nei[1]
            if len(path) < 2 || next != path[len(path)-2] { // exclude parent
                dfs(next, tail, sum + length) 
            }
        }
        last[nums[node]] = old
        path = path[:len(path)-1]
    }

    dfs(0, 0, 0)
    return []int{maxLen, max(1, minNodeCount)}
}
