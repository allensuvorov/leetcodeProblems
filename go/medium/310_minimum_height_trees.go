func findMinHeightTrees(n int, edges [][]int) []int {
    // Logic: cut leaves - till only 1 or 2 nodes left

    if n == 1 {
        return []int{0}
    }
    // Adjacency list
    adjList := make([][]int, n) 
    for _, v := range edges {
        adjList[v[0]] = append(adjList[v[0]], v[1])
        adjList[v[1]] = append(adjList[v[1]], v[0])
    }
    
    // initialize leaves 
    leaves := make([]int, 0)
    for node, v := range adjList {
        if len(v) == 1 {
            leaves = append(leaves, node)
        }
    }

    // cut (remove) leaves till 2 or less left
    for n > 2 {
        newLeaves := make([]int, 0)
        for _, leaf := range leaves {
            parent := adjList[leaf][0]

            // remove leaf from it's parent's children list
            for i, node := range adjList[parent] {
                if node == leaf {
                    adjList[parent] = append(adjList[parent][:i], adjList[parent][i+1:]...)
                    break
                }
            }
            // account for new leaves
            if len(adjList[parent]) == 1 {
                newLeaves = append(newLeaves, parent)
            }
            n--
        }
        leaves = newLeaves
    }
    return leaves
}
