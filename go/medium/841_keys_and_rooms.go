// nested DFS
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

// separate DFS
func canVisitAllRooms(rooms [][]int) bool {
    return dfs(0, rooms, map[int]bool{0:true})
}

func dfs(room int, rooms [][]int, seen map[int]bool) bool {
    seen[room] = true
    if len(seen) == len(rooms) {
        return true
    }
    for _, key := range rooms[room] {
        if !seen[key] && dfs(key, rooms, seen) {
            return true
        }
    }
    return false
}
