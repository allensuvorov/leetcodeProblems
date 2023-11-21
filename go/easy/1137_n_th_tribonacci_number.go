func tribonacci(n int) int {
    // loop solution
    a, b, c := 0, 1, 1
    abc := []int{0,1,1}
    fib := 0
    if n < 3 {
        fib = abc[n]
    }
    for i := 3; i <= n; i++ {
        fib = a + b + c
        a = b
        b = c
        c = fib
    }
    return fib
}
