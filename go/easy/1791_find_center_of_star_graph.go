func findCenter(edges [][]int) int {
    m := map[int]int{}
    for i := range edges {
        for j := range edges[i]{
            m[edges[i][j]]++
            if m[edges[i][j]] > 1 {
                return edges[i][j]
            }
        }
    }
    return 0
}
