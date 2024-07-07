func maxScore(s string) int {
    ones := 0
    for i := range s {
        if s[i] == '1' {
            ones++
        }
    }
    
    ans := 0
    zeros := 0
    for i := 0; i < len(s) - 1; i++ {
        if s[i] == '1' {
            ones--
        } else {
            zeros++
        }
        ans = max(ans, ones + zeros)
    }
    return ans
}
