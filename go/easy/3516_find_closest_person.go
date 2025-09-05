func findClosest(x int, y int, z int) int {
    if abs(z-x) < abs(z-y) {
        return 1
    }
    if abs(z-x) > abs(z-y) {
        return 2
    }
    return 0
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
