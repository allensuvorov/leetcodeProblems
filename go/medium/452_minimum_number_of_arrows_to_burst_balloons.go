func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(a, b []int) int {
        return a[1] - b[1]
    })
    arrowCount := 1
    lastEnd := points[0][1]
    for _, v := range points {
        if v[0] > lastEnd {
            arrowCount++
            lastEnd = v[1]
        }
    }
    return arrowCount
}
