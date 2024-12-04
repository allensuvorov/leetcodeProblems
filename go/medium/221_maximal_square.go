func maximalSquare(matrix [][]byte) int {
    rows, cols := len(matrix), len(matrix[0])
    res := 0
    dp := make([][]int, rows)
    for r := range dp {
        dp[r] = make([]int, cols)
    }

    for r := rows-1; r >=0; r-- {
        for c := cols-1; c >=0; c-- {
            if matrix[r][c] == '1' {
                right, down, diag := 0, 0, 0
                if c != cols-1 {
                    right = dp[r][c+1]
                }
                if r != rows-1 {
                    down = dp[r+1][c]
                } 
                if r != rows-1 && c != cols-1 {
                    diag = dp[r+1][c+1]
                }
                dp[r][c] = 1 + min(right, down, diag)
                res = max(res, dp[r][c])
            }
        }
    }
    return res*res
}


func maximalSquare(matrix [][]byte) int {
    rows, cols := len(matrix), len(matrix[0])
    res := 0
    cache := make(map[[2]int]int, rows * cols)
    
    var helper func(r, c int) int 
    helper = func(r, c int) int {
        if r >= rows || c >= cols {
            return 0
        }

        if _, ok := cache[[2]int{r,c}]; !ok {
            right := helper(r,c+1)
            down := helper(r+1,c)
            diag := helper(r+1,c+1)
            cache[[2]int{r,c}] = 0
            if matrix[r][c] == '1' {
                cache[[2]int{r,c}] = 1 + min(right, down, diag)
                res = max(res, cache[[2]int{r,c}])
            }
        }
        return cache[[2]int{r,c}]
    }

    helper(0,0)
    return res*res
}
