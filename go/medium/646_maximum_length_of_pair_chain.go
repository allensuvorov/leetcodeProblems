func findLongestChain(pairs [][]int) int {
    slices.SortFunc(pairs, func(a, b []int) int {
        if a[0] != b[0] {
            return a[0] - b[0]
        }
        return a[1] - b[1]
    })
    fmt.Println(pairs)
    
    maxLen := 1

    dp := make([]int, len(pairs))
    for i := range pairs {
        for j := 0; j < i; j++ {
            if pairs[i][0] > pairs[j][1] {
                maxLen = max(maxLen, dp[j] + 1)
            }
        }
        dp[i] = maxLen
    }
    return maxLen
}
