func uniqueOccurrences(arr []int) bool {
    occurrences := map[int]int{}
    for _, v := range arr {
        occurrences[v]++
    }

    seen := map[int]bool{} // set
    for _, v := range occurrences {
        if seen[v] {
            return false
        }
        seen[v] = true    
    }
    return true
}
