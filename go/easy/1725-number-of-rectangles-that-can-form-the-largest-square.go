func countGoodRectangles(rectangles [][]int) int {
    count := 0
    max := 0
    for _, r := range rectangles {
        shortSide := 0
        if r[0] < r[1] {
            shortSide = r[0]
        } else {
            shortSide = r[1]
        }

        if shortSide > max {
            max = shortSide
            count = 1
        } else if shortSide == max {
            count ++
        }

    }
    return count
}
