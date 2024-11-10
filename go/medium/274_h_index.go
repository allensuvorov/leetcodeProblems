func hIndex(citations []int) int {
    counts := make([]int, 1000 + 1)
    for _, v := range citations {
        counts[v]++
    }
    sum := 0
    for i := len(counts) - 1; i >= 0; i-- {
        sum += counts[i]
        if sum >= i {
            return i
        }
    }
    return 0
}
