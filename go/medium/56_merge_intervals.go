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

// different naming
func merge(intervals [][]int) [][]int {
    // sort by start
    slices.SortFunc(intervals, func(a, b []int) int{
        return a[0] - b[0]
    })
    res := make([][]int, 0)
    for _, v := range intervals {
        top := len(res) - 1 // top index
        if len(res) == 0 || v[0] > res[top][1] { // no overlap
            res = append(res, v)
        } else {
            res[top][1] = max(res[top][1], v[1]) // update end value to larger
        }
    }
    return res
}
