// editorial
func removeCoveredIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {
        if a[0] == b[0] {
            return b[1] - a[1]
        }
        return a[0] - b[0]
    })

    prevEnd := 0
    count := 0

    for _, v := range intervals {
        end := v[1]
        if prevEnd < end {
            count++
            prevEnd = end
        }
    }
    return count
}

// my solution
func removeCoveredIntervals(intervals [][]int) int {
    l, r := 0, 0
    count := 0

    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })

    for _, v := range intervals {
        if r == 0 || l != v[0] {
            if r < v[1] {
                count++
                l, r = v[0], v[1]
            }
        } else {
            r = max(r, v[1])
        }
    }
    return count
}
