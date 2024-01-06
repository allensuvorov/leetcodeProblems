func canVisitAllRooms(rooms [][]int) bool {
     visited := make([]bool, len(rooms))
     dfs(rooms, visited, 0)
     for _, ok := range visited {
         if !ok {
             return false
         }
     }
     return true
}

func dfs(rooms [][]int, visited []bool, now int){
    visited[now] = true
    for _, key := range rooms[now]{
        if !visited[key] {
            dfs(rooms, visited, key)
        }
    }
}
