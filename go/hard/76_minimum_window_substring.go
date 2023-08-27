func minWindow(s string, t string) string {
    checkList := map[byte]int{}
    
    targetMap := [127]int{}
    windowMap := [127]int{}

    for i := 0; i < len(t); i++ { // fill Target Map with letter frequencies
        targetMap[t[i]]++
        checkList[t[i]]++
    }

    // init r. r captures letters till target collection is fully captured by the window
    r := -1
    for i := 0; i < len(s); i++ {
        if targetMap[s[i]] != 0 { // fill up windowMap
            windowMap[s[i]]++
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
    
    // Second, loop through windows, find min
    l := 0
    minLen := len(s)
    ans := s

    for r < len(s) {
        // l trims tail till a must-have letter
        for ; l < len(s); l++ {
            if targetMap[s[l]] != 0 {
                if windowMap[s[l]] - targetMap[s[l]] >= 1 {
                   windowMap[s[l]]--
                } else {
                    break
                }
            }
        }
        mustHave := s[l]
        if r - l + 1 < minLen {
            minLen = r - l + 1
            ans = s[l:r+1]
        }
        
        // look for next window, update r
        r++
        for r < len(s) {
            if targetMap[s[r]] != 0 {
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
