func uniquePaths(m int, n int) int {
    row := make([]int, n)
    row[0] = 1
    for range m {
        for c := 1; c < n; c++ {
            row[c] += row[c-1]
        }
    }
    return row[len(row)-1]
}
