func assignEdgeWeights(edges [][]int) int {
    // get tree depth - d, that is the exponent
    // 2^d / 2 == 2^(d-1)
    
    t := make(map[int][]int, len(edges))
    for _, e := range edges {
        t[e[0]] = append(t[e[0]], e[1])
        t[e[1]] = append(t[e[1]], e[0])
    }
    seen := make(map[int]bool)

    maxDepth := 0
    stack := [][]int{{1,0}}
    for len(stack) > 0 {
        top := len(stack) - 1
        now := stack[top]
        stack = stack[:top] // pop
        value := now[0]
        seen[value] = true
        count := now[1]
        maxDepth = max(maxDepth, count)
        for _, nei := range t[value] {
            if !seen[nei] {
                stack = append(stack, []int{nei, count+1})
            }
        }
    }

    res := 1
    for range maxDepth {
        res = res % (1e9 + 7)
        res = res * 2
    }
    return res / 2
}
