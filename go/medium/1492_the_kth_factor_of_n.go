func kthFactor(n int, k int) int {
    
    sqrt := int(math.Sqrt(float64(n))) // 

    for i := 1; i <= sqrt; i++ {
        if n % i == 0 { // factor
            k--
        }
        if k == 0 {
            return i
        }
    }


    for i := sqrt; i > 0; i-- {
        if (i * i) == n {
            continue
        }

        if n % i == 0 {
            k--
        }

        if k == 0 {
            return n/i
        }
    }

    return -1
}
