func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })
    res := 0
    for l, r := 0, 1; r < len(intervals); r++ {
        if intervals[l][1] > intervals[r][0] {
            res++
            if intervals[l][1] > intervals[r][1] {
                l = r
            } 
        } else {
            l = r
        }
    }
    return res
}


// neetCode version
func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })
    res := 0
    prevEnd := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if start >= prevEnd {
            prevEnd = end
        } else {
            res++
            prevEnd = min(prevEnd, end)
        }
    }
    return res
}
