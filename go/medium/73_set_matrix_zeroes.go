func setZeroes(matrix [][]int)  {
    // mark zeros with MaxInt
    // Set zeros to the marked
    zeroCal0 := false
    zeroRow0 := false

    for c := range matrix[0] {
        if matrix[0][c] == 0 {
            zeroRow0 = true
        }
    }

    for r := range matrix {
        if matrix[r][0] == 0 {
            zeroCal0 = true
        }
    }

    for r := 1; r < len(matrix); r++ {
        for c := 1; c < len(matrix[0]); c++ {
            if matrix[r][c] == 0 {
                matrix[r][0] = 0
                matrix[0][c] = 0
            }
        }
    }

    for c := 1; c < len(matrix[0]); c++ {
        if matrix[0][c] == 0 {
            for r := range matrix{
                matrix[r][c] = 0
            }
        }
    }

    for r := 1; r < len(matrix); r++ {
        if matrix[r][0] == 0 {
            for c := range matrix[0] {
                matrix[r][c] = 0
            }
        }
    }
    if zeroCal0 {
        for r := range matrix {
            matrix[r][0] = 0
        }
    }

    if zeroRow0 {
        for c := range matrix[0] {
            matrix[0][c] = 0
        }
    }
}
