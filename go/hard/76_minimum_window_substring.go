func minWindow(s string, t string) string {
    tCount, window := [127]int{}, [127]int{}
    for i := 0; i < len(t); i++ {
        tCount[t[i]]++
    }
    
    have, need := 0, 0
    for _, v := range tCount {
        if v != 0 {
            need++
        }
    }
    
    res, resLen := []int{-1,-1}, len(s) + 10
    l := 0
    for r := range s {
        if tCount[s[r]] != 0{
            window[s[r]]++
            if window[s[r]] == tCount[s[r]] {
                have++
            }
        }

        for have == need {
            // update results
            if (r - l + 1) < resLen {
                res = []int{l, r}
                resLen = r - l + 1
            }
            // pop from the left of our window
            window[s[l]]--
            if tCount[s[l]] != 0 && window[s[l]] < tCount[s[l]] {
                have--
            }
            l++
        }
    }
    if resLen == len(s) + 10 {
        return ""
    }
    return s[res[0]:(res[1] + 1)]
}
