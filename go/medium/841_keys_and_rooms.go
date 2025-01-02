func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visitedCount := 0
	var dfs func(room int) 
    dfs = func(room int) {
        if visited[room] {
            return
        }
        visited[room] = true
        visitedCount++
        for _, subRoom := range rooms[room] {
            dfs(subRoom)
        }
    }
    dfs(0)
    return visitedCount == len(rooms)
}
