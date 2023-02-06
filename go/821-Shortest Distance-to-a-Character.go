func shortestToChar(s string, c byte) []int {
    prev := -100000
    ans := make([]int, len(s))
    
    for i := range s {
        if s[i] == c {
            prev = i
        }
        ans[i] = i - prev
    }
    
    prev = math.MaxInt
    for i := len(s)-1; i >=0; i-- {
        if s[i] == c {
            prev = i
        }
        if ans[i] > prev - i {
            ans[i] = prev - i 
        }
    }
    return ans
}
