func minOperations(grid [][]int, x int) int {
    // is it unifiable?
    target := grid[0][0]
    for r := range grid {
        for c := range grid[0] {
            if abs(grid[r][c] - target) % x != 0 {
                return -1
            }
        }
    }

    sum := 0
    for r := range grid {
        for c := range grid[0] {
            sum += grid[r][c]
        }
    }

    average := sum / ( len(grid) * len(grid[0]) )

    // closest product of x + 
    diff := abs(target-average)

    lowerTarget := diff - diff % x
    upperTarget := lowerTarget + x

    opsToReachLower := 0
    opsToReachUpper := 0

    for r := range grid {
        for c := range grid[0] {
            opsToReachLower += abs(grid[r][c] - lowerTarget)/x
            opsToReachUpper += abs(grid[r][c] - upperTarget)/x
        }
    } 

    return min(opsToReachLower, opsToReachUpper)
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
