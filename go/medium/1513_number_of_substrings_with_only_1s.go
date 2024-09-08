func numSub(s string) int {
    ans := 0
    curSum := 0
    count := 0
    for _, v := range s {
        if v == '1' {
            count++
            curSum += count
            curSum %= 1e9 + 7
        } else {   
            ans += curSum 
            ans %= 1e9 + 7
            count = 0
            curSum = 0
        }
    }
    ans += curSum 
    return ans
}
