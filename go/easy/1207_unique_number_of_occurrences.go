func uniqueOccurrences(arr []int) bool {

    numCounts := make(map[int]int)
    for _, v := range arr {
        numCounts[v]++
    }

    countCounts := make(map[int]int)
    for _, v := range numCounts {
        if countCounts[v] > 0 {
            return false
        }
        countCounts[v]++
    }

    return true
}
