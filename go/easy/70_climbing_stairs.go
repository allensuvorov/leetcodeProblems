func climbStairs(n int) int {
    one, two := 1, 1
    for i := 1; i < n; i++ {
        temp := one
        one += two
        two = temp
    }
    return one
}
