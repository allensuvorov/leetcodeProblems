func xorOperation(n int, start int) int {
    xSum := 0
    for i := 0; i < n; i ++ {
        xSum = xSum ^ (start + 2 * i)
    }
    return xSum
}
