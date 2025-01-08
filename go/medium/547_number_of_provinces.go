func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    provCount := 0
    for i := range n {
        if isConnected[i][i] == 1 {
            provCount++
            bfs(isConnected, i)
        }
    }
    return provCount
}

func bfs(isConnected [][]int, now int) {
    q := []int{now}
    for len(q) > 0 {
        now := q[0]
        q = q[1:]
        isConnected[now][now] = 2
        for nei := range isConnected {
            if isConnected[now][nei] == 1 && isConnected[nei][nei] == 1 {
                q = append(q, nei)
            }
        }
    }
}
