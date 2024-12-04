func isInterleave(s1 string, s2 string, s3 string) bool {
    l1, l2, l3 := len(s1), len(s2), len(s3)
    if l1 + l2 != l3 {
        return false
    }

    dp := make([][]bool, l1 + 1)
    for i := range dp {
        dp[i] = make([]bool, l2 + 1)
    }
    dp[l1][l2] = true

    for i := l1; i >= 0; i-- {
        for j := l2; j >= 0; j-- {
            if i < l1 && s1[i] == s3[i + j] && dp[i + 1][j] {
                dp[i][j] = true
            }
            if j < l2 && s2[j] == s3[i + j] && dp[i][j + 1] {
                dp[i][j] = true
            }
        }
    }
    return dp[0][0]
}
