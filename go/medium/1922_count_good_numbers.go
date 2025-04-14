func countGoodNumbers(n int64) int {
    mod := int64(1e9 + 7)

    quickMul := func(x, y int64) int64 {
        res := int64(1)
        mul := x
        for y > 0 {
            if y%2 == 1 {
                res = res * mul % mod
            }
            mul = mul * mul % mod
            y /= 2
        }
        return res
    }

    return int( quickMul(5, (n+1)/2) * quickMul(4, n/2) % mod)
}
