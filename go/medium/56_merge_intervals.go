func merge(intervals [][]int) [][]int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })

    merged := make([][]int, 0)
    for _, interval := range intervals {
        last := len(merged) - 1
        if len(merged) == 0 || merged[last][1] < interval[0] {
            merged = append(merged, interval)
        } else {
            merged[last][1] = max(merged[last][1], interval[1])
        }
    }
    return merged
}
