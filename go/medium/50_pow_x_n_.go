func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    
    if x == 0 {
        return 0
    }
    
    if n == 1 {
        return x
    }

    if x == 1 {
        return 1
    }

    negExp := false
    if n < 0 {
        negExp = true
        n = -n
    }

    a := x
    ans := float64(1)

    for n > 0 {
        if n % 2 == 1 {
            ans *= a
        }
        a *= a
        n /= 2
    }

    if negExp {
        ans = 1/ans
    }

    return ans
}
