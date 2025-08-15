func isPowerOfFour(n int) bool {
    sqrt := math.Sqrt(float64(n))
    if sqrt != float64(int(sqrt)) {
        return false
    }
    intSqRt := int(sqrt)

    return n > 0 && intSqRt & (intSqRt - 1) == 0
}
