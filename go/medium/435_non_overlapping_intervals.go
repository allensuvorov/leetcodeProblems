func eraseOverlapIntervals(intervals [][]int) int {
    // preprocessing - sort
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })

    lastEnd := intervals[0][1]
    result := 0
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        // detect overlap, update lastEnd
        if start < lastEnd {
            result++
            lastEnd = min(lastEnd, end)
        } else {
            lastEnd = end
        }
    }
    return result
}
