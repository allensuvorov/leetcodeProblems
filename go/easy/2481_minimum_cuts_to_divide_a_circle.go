func numberOfCuts(n int) int {
    if n > 1 {
        if n % 2 == 0 {
            return n / 2
        } else {
            return n
        }
    }
    return 0
}
