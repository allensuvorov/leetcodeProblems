func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(a, b []int) int {
        return a[1]-b[1] 
    })
    lastEnd := math.MinInt
    arrowCount := 0
    for _, v := range points {
        if lastEnd < v[0] {
            arrowCount++
            lastEnd = v[1]
        }
    }
    return arrowCount
}
