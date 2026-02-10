func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0] // sort by start
    })
    res := 0
    reach := math.MinInt
    for _, v := range intervals{
        if v[0] < reach { // if there is overlap
            res++
            reach = min(reach, v[1]) // keep the shorter reach
        } else {
            reach = v[1]
        }
    }
    return res
}
