func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	// look for a zero, save row and col
	cacheRow, cacheCol := -1, -1
	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 0 {
				cacheRow = r
				cacheCol = c
			}
		}
	}

    // if no zero, return
	if cacheRow == -1 {
		return
	}

	// use those cache fields as memory
	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 0 {
				matrix[r][cacheCol] = 0
				matrix[cacheRow][c] = 0
			}
		}
	}

	// zero target rows, except cache row
	for r := range rows {
		if r != cacheRow && matrix[r][cacheCol] == 0 {
			for c := range cols {
				matrix[r][c] = 0
			}
		}
	}
    
	// zero target cols, except cache col
	for c := range cols {
		if c != cacheCol && matrix[cacheRow][c] == 0 {
			for r := range rows {
				matrix[r][c] = 0
			}
		}
	}
    
    // zero cache row
    for c := range cols {
        matrix[cacheRow][c] = 0
    }
    
    // zero cache col
    for r := range rows {
        matrix[r][cacheCol] = 0
    }
}
