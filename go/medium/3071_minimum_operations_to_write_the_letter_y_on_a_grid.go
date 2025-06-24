func minimumOperationsToWriteY(grid [][]int) int {
    // count frequencies
    Y := make([]int,3) //
    X := make([]int,3) // non Y
    n := len(grid)
    for r := range n {
        for c := range n {
            // check if letter belongs to Y
            lY := r == c && r <= n/2 && c <= n/2
            rY := r + c == n - 1 && r <= n/2 && c >= n/2
            mY := c == n / 2 && r >= n/2
            if lY || rY || mY {
                Y[grid[r][c]]++
            } else {
                X[grid[r][c]]++
            }
        }
    }
    res := n*n
    for a := range 3 {
        res = min(res, operationsToWriteY(Y, X, a))
    }
    return res
}

func operationsToWriteY(Y, X []int, a int) int {
    // write Y with number a
    writeY := change(Y, a)
    // write X, the background, with other than a
    minToWriteX := math.MaxInt
    for b := range X {
        if b != a {
            minToWriteX = min(minToWriteX, change(X, b))
        }
    }
    return writeY + minToWriteX
}

func change(nums []int, a int) int {
    res := 0
    for b := range nums {
        if b != a {
            res += nums[b]  
        }
    }
    return res
}
