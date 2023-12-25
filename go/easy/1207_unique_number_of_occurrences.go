func uniqueOccurrences(arr []int) bool {
    fm := make(map[int]int)
    for _, v := range arr {
        fm[v]++
    }
    fs := make(map[int]bool)
    for _, v := range fm {
        if fs[v] {
            return false
        }
        fs[v] = true
    }
    return true
}
