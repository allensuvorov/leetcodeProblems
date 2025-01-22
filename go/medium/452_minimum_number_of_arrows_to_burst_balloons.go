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

// redo
func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(a,b []int)int{
        return a[1]-b[1]
    })
    arrowCount := 0
    prevEnd := math.MinInt
    for _, v := range points {
        start, end := v[0], v[1]
        if prevEnd < start {
            arrowCount++
            prevEnd = end
        }
    }
    return arrowCount
}

// [[3,9],[7,12],[3,8],[6,8],[9,10],[2,9],[0,9],[3,9],[0,6],[2,8]]

// 0-1-2-3-4-5-6-7-8-9-10-11-12-13-14-15-16-17-18

// 0-----------6
//       3---------8
//             6---8
//     2-----------8
//       3-----------9
//     2-------------9
// 0-----------------9
//       3-----------9
//                   9-10
//               7-----------12



// Input: points = [[10,16],[2,8],[1,6],[7,12]]

// 1-2-3-4-5-6-7-8-9-10-11-12-13-14-15-16-17-18
// 1---------6
//   2-----------8
//             7-----------12
//                   10----------------16

//  1           2             
