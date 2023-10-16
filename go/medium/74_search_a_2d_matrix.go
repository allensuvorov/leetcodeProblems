func searchMatrix(matrix [][]int, target int) bool {
    // find the row (smallest that's with first element lower than target)
    l := 0
    r := len(matrix) - 1
    for l < r {
        m := l + (r - l + 1)/2
        if matrix[m][0] <= target {
            l = m
        } else {
            r = m - 1
        }
    }
    
    // look within that row
    row := l
    l = 0
    r = len(matrix[0]) - 1
    for l <= r {
        m := l + (r - l)/2
        guess := matrix[row][m]
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
