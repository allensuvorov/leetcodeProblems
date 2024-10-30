func findMinArrowShots(points [][]int) int {
    // sort (inc)
    slices.SortFunc(points, func(a, b []int) int {
        return a[1] - b[1]
    })

    arrows := 1
    // at each arrow track earliest END, then look for next START > END
    for l, r:= 0, 1; r < len(points); r++ {
        if points[r][0] > points[l][1] {
            arrows++
            l = r
        }
    }
    return arrows

}
