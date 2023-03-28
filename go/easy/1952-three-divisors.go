func isThree(n int) bool {
    d := 2
    for i := 2; i < n; i ++ {
        if n % i == 0 {
            d++
        }
        if d > 3 {
            return false
        }
    }
    return d == 3
}
