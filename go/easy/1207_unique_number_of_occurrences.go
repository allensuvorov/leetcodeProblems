func uniqueOccurrences(arr []int) bool {
    // get counts
    counts := make(map[int]int)
    
    for _, v := range arr {
        counts[v]++
    }

    
    // check counts for repeating a value
    hasCount := make(map[int]bool)

    for _, v := range counts {
        if hasCount[v] {
            return false
        }
        hasCount[v] = true
    }
    return true
}
