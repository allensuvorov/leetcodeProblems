func isPowerOfTwo(n int) bool {
    return n > 0 && n & (n-1) == 0
}

//
func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }

    return 0 == n & (n - 1) // only a number like 10000 will lose the most significant bit, when we bitwise AND it with itself after subtracting 1.
}
