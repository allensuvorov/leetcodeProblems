func tribonacci(n int) int {
    if n < 2 {
        return n
    }

    a, b, c := 0, 1, 1
    for range n - 2 {
        temp := c
        c = a + b + c
        a = b
        b = temp
    }
    return c
}
