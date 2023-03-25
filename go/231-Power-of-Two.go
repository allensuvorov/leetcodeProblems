func isPowerOfTwo(n int) bool {
    return int(math.Pow(math.Sqrt(float64(n)), 2)) == n
}
