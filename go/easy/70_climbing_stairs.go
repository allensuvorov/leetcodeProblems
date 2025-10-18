func climbStairs(n int) int {
    one, two := 1, 1
    for i := 1; i < n; i++ {
        temp := one
        one += two
        two = temp
    }
    return one
}

func climbStairs(n int) int {
    a, b := 1, 2
    for range n - 1 {
        a, b = b, a + b
    }
    return a
}
