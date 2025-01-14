func numTilings(n int) int {
    // try to find the pattern
    // use the pattern
    if n < 3 {
        return n
    }
    a, b, c := 1, 2, 5
    for range n - 3 {
        a, b, c = b, c, 2*c + a
        c %= 1e9 + 7
    }
    return c
}
// 1 - 1
// 2 - 2
// 3 - 5
// 4 - 11 (2*5 + 1)
// 5 - 24 (2*11 + 2)
