func countBits(n int) []int {
    ans := make([]int, n + 1)
    for i := 1; i <= n; i++ {        
        if i & 1 == 0 { // even
            ans[i] = ans[i/2]
        } else { // odd
            ans[i] = ans[i-1] + 1
        }
    }
    return ans
}
