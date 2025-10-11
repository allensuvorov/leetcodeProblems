func maxRectangleArea(points [][]int) int {
    maxArea := -1

    // find 4 valid points
    for i1 := range points {
        a := points[i1]
        for i2 := range points {
            b := points[i2]
            if i1 != i2 && b[1] == a[1] {
                for i3 := range points {
                    c := points[i3]
                    if i3 != i2 && i3 != i1 && c[0] == b[0] {
                        for i4 := 0; i4 < len(points); i4++ {
                            d := points[i4]
                            if i4!= i3 && i4 != i2 && i4 != i1 && d[0] == a[0] && d[1] == c[1] {
                                // check overlaps
                                containsOtherPoint := false
                                for j := range points {
                                    if j != i1 && j != i2 && j != i3 && j != i4 {
                                        e := points[j]
                                        if e[0] >= min(a[0], b[0]) && 
                                           e[0] <= max(a[0], b[0]) && 
                                           e[1] >= min(d[1], a[1]) && 
                                           e[1] <= max(d[1], a[1]) {
                                            containsOtherPoint = true
                                            break
                                        }
                                    }
                                }
                                if !containsOtherPoint {
                                    area := abs(b[0] - a[0]) * abs(a[1] - d[1])
                                    maxArea = max(maxArea, area)
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    return maxArea
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
