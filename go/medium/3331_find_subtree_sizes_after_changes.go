func findSubtreeSizes(parent []int, s string) []int {
    n := len(parent) 
    g := make([][]int, n) // create top->down graph
    for i := 1; i < n; i++ {
        p := parent[i]
        g[p] = append(g[p], i)
    }

    size := make([]int, n)
    ancestor := [26]int{} // ancestors

    for i := range ancestor {
        ancestor[i] = -1
    }
    var dfs func(x int)
    dfs = func(x int) {
        size[x] = 1 // count the node itself
        old := ancestor[s[x]-'a'] // save old ancestor
        ancestor[s[x]-'a'] = x
        for _, y := range g[x] { // y is the child
            dfs(y)
            anc := x // be default x is the new parent
            if v := ancestor[s[y]-'a']; v != -1 {
                anc = v
            }
            size[anc] += size[y]
        }
        ancestor[s[x]-'a'] = old
    }
    dfs(0)
    return size
}
