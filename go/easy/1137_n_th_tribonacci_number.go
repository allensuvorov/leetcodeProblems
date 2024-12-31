func tribonacci(n int) int {
    if n < 2 {
        return n
    }
    if n == 2 {
        return 1
    }

    a, b, c := 0, 1, 1

    for range n - 2 {
        a, b, c = b, c, a + b + c
    }
    return c
}
