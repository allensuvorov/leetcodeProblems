func myPow(x float64, n int) float64 {
    if x == 0 {
        return 0
    }
    
    if n == 0 {
        return 1
    }
    
    if x == 1 {
        return 1
    }

    ans := myPow(x, abs(n) / 2)
    ans *= ans
    
    if abs(n) % 2 == 1 {
        ans *= x
    }

    if n < 0 {
        ans = 1 / ans
    }

    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
