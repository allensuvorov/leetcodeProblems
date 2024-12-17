func maxPoints(points [][]int) int {
    n := len(points)

    res := 0
    for i := range n{
        slopes := make(map[float64]int)
        for j := range n {
            m := slope(points[i], points[j])
            slopes[m]++
            res = max(res, slopes[m])
        }
    }
    return res + 1
}


func slope(point1, point2 []int) float64 {
    x1, y1 := point1[0], point1[1]
    x2, y2 := point2[0], point2[1]

    slope := float64(y1 - y2) / float64(x1 - x2) // slope
    return slope
}
