func searchMatrix(matrix [][]int, target int) bool {
    l := 0
    r := len(matrix) * len(matrix[0]) - 1
    for l <= r {
        m := l + (r - l)/2
        // convert number into coordinates
        i := m / len(matrix[0])
        j := m % len(matrix[0])
        guess := matrix[i][j]
        switch {
        case guess == target:
            return true
        case guess > target:
            r = m - 1
        case guess < target:
            l = m + 1
        }
    }
    return false
}
