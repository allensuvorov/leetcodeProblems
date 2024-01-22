func fib(n int) int {
    a, b := 0, 1
    for i := 1; i <= n; i++ {
        a, b = b, a+b
    }
    return a
}
