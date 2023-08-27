func minWindow(s string, t string) string {
    // fill Target Map with letter frequencies
    tm := map[byte]int{}
    for _, v := range t {
        tm[v]++
    }
    
    // look for the first window, initialise l and r 

    // init l
    l := -1
    for i, v := range s {
        if _, ok := tm[v]; ok{
            l = i
            break
        }
    }
    if l == -1 {
        return ""
    }
    // init r
    // r captures letters till target collection is fully captured
    r := -1
    for i := l; i < len(s); i++ {
        tm[s[i]]--
        if tm[s[i]] == 0
        delete(tm, s[i])
        if len(tm) == 0 {
            r = i
            break
        }
    }
    if r == -1 {
        return ""
    }
    
    
    wm(s[r])++
    wm := map[byte]int{}
    


    ans := ""
    minLen := len(t)
    return ans
}
