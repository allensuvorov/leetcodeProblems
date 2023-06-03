func findColumnWidth(grid [][]int) []int {
    max := 0
    ans := make([]int, len(grid[0]))

    for c := range grid[0] {
        for r := range grid {
            w := intWidth(grid[r][c])
            if max < w {
                max = w
            }
        }
        ans[c] = max
        max = 0
    }
    return ans
}

func intWidth(n int) int {
    digCount := 0
    if n < 0 {
        digCount++
    }
    if n == 0 {
        return 1
    }
    for n != 0 {
        n /= 10
        digCount++
    }
    return digCount
}

//   c c c
// r
// r
// r
