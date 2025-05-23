func numTrees(n int) int {
    dp := make([]int, n + 1)
    dp[0] = 1 // to support combinations with zero nodes
    
    for nodeCount := 1; nodeCount <= n; nodeCount++ {
        totalTreeCount := 0
        for root := 1; root <= nodeCount; root++ {
            leftSubTreeSize := root - 1
            rightSubTreeSize := nodeCount - root
            totalTreeCount += dp[leftSubTreeSize] * dp[rightSubTreeSize]
        }
        dp[nodeCount] = totalTreeCount
    }
    
    return dp[n]
}
