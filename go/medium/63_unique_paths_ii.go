func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		obstacleGrid[0][0] = 0
	} else {
		obstacleGrid[0][0] = 1
	}

	for c := 1; c < cols; c++ {
		if obstacleGrid[0][c] == 1 {
			obstacleGrid[0][c] = 0
		} else {
			obstacleGrid[0][c] = obstacleGrid[0][c-1]
		}
	}

	for r := 1; r < rows; r++ {
		if obstacleGrid[r][0] == 1 {
			obstacleGrid[r][0] = 0
		} else {
			obstacleGrid[r][0] = obstacleGrid[r-1][0]
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if obstacleGrid[r][c] == 1 {
				obstacleGrid[r][c] = 0
			} else {
				obstacleGrid[r][c] = obstacleGrid[r-1][c] + obstacleGrid[r][c-1]
			}
		}
	}

	return obstacleGrid[rows-1][cols-1]
}
