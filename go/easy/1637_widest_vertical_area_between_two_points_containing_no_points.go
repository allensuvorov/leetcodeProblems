import "sort"
func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0]})
    ans, prev := 0, points[0][0]
    for i := range points {
        ans = max(points[i][0] - prev, ans)
        prev = points[i][0]
    }
    return ans
}
