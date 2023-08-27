func minWindow(s string, t string) string {
    targetMap := map[byte]int{}
    checkList := map[byte]int{}
    windowMap := map[byte]int{}

    for i := 0; i < len(t); i++ { // fill Target Map with letter frequencies
        targetMap[t[i]]++
        checkList[t[i]]++
    }
    // look for the initial window l and r 
    // first, init r. r captures letters till target collection is fully captured by the window
    r := -1
    for i := 0; i < len(s); i++ {
        if _, ok := targetMap[s[i]]; ok { // fill up windowMap
            windowMap[s[i]]++
            fmt.Println(checkList)
            if checkList[s[i]] != 0 {
                checkList[s[i]]-- // to track progress
            } 
            if checkList[s[i]] == 0 {
                delete(checkList, s[i])
            }
        }
        if len(checkList) == 0 {
            r = i
            break
        }
    }
    if r == -1 {
        return ""
    }
    
    l := 0
    minLen := len(s)
    ans := s

    for r < len(s) {
        // l trims tail till a must-have letter
        for ; l < len(s); l++ {
            if _, ok := targetMap[s[l]]; ok {
                if windowMap[s[l]] - targetMap[s[l]] >= 1 {
                   windowMap[s[l]]--
                } else {
                    break
                }
            }
        }
        mustHave := s[l]
        fmt.Println(l, r)
        if r - l + 1 < minLen {
            minLen = r - l + 1
            ans = s[l:r+1]
        }
        
        // look for next window, update r
        r++
        for r < len(s) {
            if _, ok := targetMap[s[r]]; ok {
                windowMap[s[r]]++
                if s[r] == mustHave {
                    break
                }
            }
            r++
        }
    }
    return ans
}
