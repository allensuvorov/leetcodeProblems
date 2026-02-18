func findMinArrowShots(points [][]int) int {
    // soft by end
    slices.SortFunc(points, func(a, b []int) int {
        return a[1] - b[1]
    })
    // count arrows
    res := 0
    lastEnd := math.MinInt
    for _, v := range points {
        if lastEnd < v[0] {
            res++
            lastEnd = v[1]
        }
    }
    return res
}
