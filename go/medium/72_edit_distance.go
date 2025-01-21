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
			}

			dp[r][c] = prevMinVal + 1

			if word1[c-1] == word2[r-1] {
				dp[r][c] = dp[r-1][c-1]
			}

		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// redo
func minDistance(word1 string, word2 string) int {
    if len(word1) == 0 || len(word2) == 0 {
        return max(len(word1),len(word2))
    }
    dp := make([][]int, len(word2))
    for r := range word2 {
        dp[r] = make([]int, len(word1))
        for c := range word1 {
            l, t, d := r+1, c+1, 0
            
            if r > 0 && c == 0{
                t = dp[r-1][c]
                d = r
            }

            if c > 0 && r == 0{
                l = dp[r][c-1]
                d = c
            }
            
            if r > 0 && c > 0 {
                l = dp[r][c-1]
                t = dp[r-1][c]
                d = dp[r-1][c-1]
            }

            if word1[c] != word2[r] {
                dp[r][c] = 1 + min(l, t, d)
            } else {
                dp[r][c] = d
            }
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}
