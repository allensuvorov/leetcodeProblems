func minSteps(n int) int {
    ans := 0
    for d := 2; n > 1; d++ {
        for n % d == 0 {
            ans += d
            n /= d
        }
    } 
    return ans
}
