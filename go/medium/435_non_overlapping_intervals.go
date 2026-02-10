func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0] // sort by start
    })
    res := 0
    lastEnd := math.MinInt
    for _, v := range intervals{
        if v[0] < lastEnd { // if there is overlap
            res++
            lastEnd = min(lastEnd, v[1]) // keep the shorter lastEnd
        } else {
            lastEnd = v[1]
        }
    }
    return res
}
