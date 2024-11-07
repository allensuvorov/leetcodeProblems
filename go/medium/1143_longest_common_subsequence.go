func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int, len(text2))
    for r := range text2 {
        dp[r] = make([]int, len(text1))
        for c := range text1 {
            if text2[r] == text1[c] {
                if c > 0 && r > 0 {
                    dp[r][c] = dp[r-1][c-1] + 1
                } else {
                    dp[r][c] = 1
                }
            } else {
                top := 0
                if r > 0 {
                    top = dp[r-1][c]
                }
                left := 0
                if c > 0 {
                    left = dp[r][c-1]
                }
                dp[r][c] = max(top, left)
            }   
        }
    }
    return dp[len(dp) - 1][len(dp[0])-1]
}
