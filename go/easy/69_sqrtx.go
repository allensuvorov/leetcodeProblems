func mySqrt(x int) int {
    l, r := 0, x
    for l < r {
        m := (r + l)/2 + 1
        if m*m > x {
            r = m - 1
        } else {
            l = m
        }
    }
    return r
}
