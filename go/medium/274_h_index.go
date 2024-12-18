func hIndex(citations []int) int {
    counts := make([]int, 1000 + 1)
    for _, v := range citations {
        counts[v]++
    }
    c := 1000
    for sum := counts[c]; c > sum; sum += counts[c] {
        c--
    }
    return c
}
