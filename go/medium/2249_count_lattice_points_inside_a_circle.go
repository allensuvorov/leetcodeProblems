func countLatticePoints(circles [][]int) int {
    res := 0
    for x := 0; x <= 200; x++ {
        for y := 0; y <= 200; y++ {
            if isPointInCircles(circles, x, y) {
                res++
            }
        }
    }
    return res
}

func isPointInCircles(circles [][]int, x, y int) bool {
    for _, v := range circles {
        side1 := x - v[0]
        side2 := y - v[1]
        if side1*side1 + side2*side2 <= v[2] * v[2] {
            return true
        }
    }
    return false
}
