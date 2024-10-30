func eraseOverlapIntervals(intervals [][]int) int {
    // sort by starts (inc)
    slices.SortFunc(intervals, func(a, b []int) int{
        return a[0] - b[0]
    })

    skips := 0
    for l, r := 0, 1; r < len(intervals); r++ {
        // detect overlap
        if intervals[r][0] < intervals[l][1] {
            skips++
            // skip interval that stretchs further
            if intervals[l][1] >= intervals[r][1] {
                l = r
            }
        } else {
            l = r
        }
    }
    return skips
}
