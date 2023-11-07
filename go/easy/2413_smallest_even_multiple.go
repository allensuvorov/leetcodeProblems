func smallestEvenMultiple(n int) int {
    if n%2 == 0 {
        return n
    }
    return n*2
}

// 1 - 2 (x2)
// 2 - 2
// 3 - 6 (x2)
// 4 - 4
// 5 - 10 (x2)
// 6 - 6
// 7 - 14 (x2)
