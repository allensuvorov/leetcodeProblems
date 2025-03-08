func closestPrimes(left int, right int) []int {
    primes := []int{0:1, 1:1, 1e6:0} // 0 - prime, 1 - not a prime
    for i := 2; i <= right; i++ {
        if primes[i] == 0 {
            for j := i + i; j <= right; j += i {
                primes[j] = 1
            }
        }
    }
    
    minDiff := math.MaxInt
    num1, num2 := -1, -1
    res := []int{num1, num2}
    for i := left; i <= right; i++ {
        if primes[i] == 0 {
            num1, num2 = num2, i
            if num1 > 0 && num2 > 0 && (num2 - num1) < minDiff{
                minDiff = num2 - num1
                res = []int{num1, num2}
            }
        }
    }

    return res
}
