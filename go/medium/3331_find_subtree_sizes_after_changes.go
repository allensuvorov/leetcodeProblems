func findSubtreeSizes(parent []int, s string) []int {
    // First - update parents
    // to go top->down, need to create adjList - children
    children := make([][]int, len(parent))
    for i := range parent {
        if i != 0 {
            children[parent[i]] = append(children[parent[i]], i)
        }
    }
    ancestors := make(map[byte][]int) // map of stacks of ansestors

    var updateParent func(number int)
    updateParent = func(number int) {
        // update parent if needed
        if ancestors[s[number]] != nil && len(ancestors[s[number]]) > 0{
            top := len(ancestors[s[number]]) - 1
            parent[number] = ancestors[s[number]][top]
        }
        ancestors[s[number]] = append(ancestors[s[number]], number)
        for _, child := range children[number] {
            updateParent(child)
        }
        top := len(ancestors[s[number]]) - 1
        ancestors[s[number]] = ancestors[s[number]][:top]
    }

    updateParent(0)
    // Second - count sizes
    // update adjList children
    children = make([][]int, len(parent))
    for i := range parent {
        if i != 0 {
            children[parent[i]] = append(children[parent[i]], i)
        }
    }

    res := make([]int, len(parent))

    var countSize func(number int) int
    countSize = func(number int) int {
        size := 1
        for _, child := range children[number] {
            size += countSize(child)
        }
        res[number] = size
        return size
    }
    countSize(0)
    return res
}
