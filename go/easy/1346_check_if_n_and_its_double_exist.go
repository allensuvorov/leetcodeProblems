func checkIfExist(arr []int) bool {
    seen := make(map[int]bool)
    for _, v := range arr {
        if seen[v*2] || (v % 2 == 0 && seen[v/2]) {
            return true
        }
        seen[v] = true
    }
    return false
}
