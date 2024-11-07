func numTilings(n int) int {
    if n < 3 {
        return n
    }

    a, b, c := 1, 2, 5

    for range n - 3 {
        a, b, c = b, c, (c*2 + a)
        c %= (1e9 + 7)
    }
    return c
}
