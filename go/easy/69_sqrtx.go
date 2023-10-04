func mySqrt(x int) int {
    l, r := 0, x
    for l < r {
        m := (l + r + 1)/2    
        if m*m <= x {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}
