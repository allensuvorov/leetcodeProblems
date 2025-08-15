func isPowerOfFour(n int) bool {
    sqrt := math.Sqrt(float64(n))
    if sqrt != float64(int(sqrt)) {
        return false
    }
    intSqRt := int(sqrt)

    return n > 0 && intSqRt & (intSqRt - 1) == 0
}

// Big O (log(n))
func isPowerOfFour(n int) bool {
    for n % 4 == 0 {
        n = n / 4
    }
    return n == 1
}

// Super cool bitwise solution, that uses base 16 and 

func isPowerOfFour(n int) bool {
    return n > 0 && n & (n - 1) == 0 && (n&0xaaaaaaaa == 0)
}
