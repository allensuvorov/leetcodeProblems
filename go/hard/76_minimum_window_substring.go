func minWindow(s string, t string) string {
    tCount, window := map[byte]int{}, [127]int{}

    for i := 0; i < len(t); i++ {
        tCount[t[i]]++
    }

    have, need := 0, len(tCount)
    res, resLen := []int{-1,-1}, len(s) + 10
    l := 0
    for r := range s {
        c := s[r]
        window[c]++

        if tCount[c] != 0 && window[c] == tCount[c]{
            have++
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
