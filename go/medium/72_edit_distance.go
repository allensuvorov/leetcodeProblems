func minDistance(word1 string, word2 string) int {


	dp := make([][]int, len(word2) + 1)
    dp[0] = make([]int, len(word1) + 1)
    for j := range dp[0] {
        dp[0][j] = j
    }

	for r := 1; r < len(word2) + 1; r++ {
		dp[r] = make([]int, len(word1) + 1)
		dp[r][0] = r
        for c := 1; c < len(word1) + 1; c++{
			prevMinVal := 0
			if c > 0 && r > 0 {
				prevMinVal = min(dp[r-1][c-1], dp[r-1][c], dp[r][c-1])
			} else if r > 0 {
				prevMinVal = dp[r-1][c]
			} else if c > 0 {
				prevMinVal = dp[r][c-1]
			}

			val := 1
			if word1[c-1] == word2[r-1] {
				val = 0
			}

			dp[r][c] = prevMinVal + val
		}
	}

    fmt.Print("     ")
    for i := range word1 {
        fmt.Printf("%c  ", word1[i])
    }
    fmt.Println()

	for r := range len(word2) + 1 { 
		if r > 0 {
            fmt.Printf("%c ", word2[r-1])
        } else {
            fmt.Print("  ")
        }
		for c := range len(word1) + 1{
			fmt.Printf("%02d ", dp[r][c])
		}
		fmt.Println()
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
