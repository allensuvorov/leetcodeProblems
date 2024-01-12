func minReorder(n int, connections [][]int) int {
    adjacent := make([][]int, n)
    conMap := make([]map[int]bool, n)

    for _, edge := range connections {
        adjacent[edge[0]] = append(adjacent[edge[0]], edge[1])
        adjacent[edge[1]] = append(adjacent[edge[1]], edge[0])
        if conMap[edge[0]] == nil {
            conMap[edge[0]] = map[int]bool{edge[1]:true}
        } else {
            conMap[edge[0]][edge[1]] = true
        }
    }

    count := 0
    visited := make([]bool, n)

    var dfs func(now int)
    dfs = func(now int){
        if !visited[now] {
            visited[now] = true
            for _, city := range adjacent[now] {
                if !visited[city] {
                    if !conMap[city][now] {
                        count++
                    }
                    dfs(city)
                }
            }
        }
    }

    dfs(0)
    return count
}
