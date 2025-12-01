// iterative DFS with a stack
func canVisitAllRooms(rooms [][]int) bool {
    todo := []int{0}
    visited := make(map[int]bool)
    for len(todo) > 0 {
        now := todo[len(todo)-1]
        todo = todo[:len(todo)-1]
        visited[now] = true
        for _, nei := range rooms[now] {
            if !visited[nei] {
                todo = append(todo, nei)
            }
        }
    }
    return len(visited) == len(rooms) 
}

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
