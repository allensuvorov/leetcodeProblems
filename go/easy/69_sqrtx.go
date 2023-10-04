func mySqrt(x int) int {
    z := 1 
    for z*z <= x {
        if z*z == x {
            return z
        }
        z++
    }
    return z - 1
}
