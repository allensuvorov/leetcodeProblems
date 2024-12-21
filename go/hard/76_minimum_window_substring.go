func minWindow(s string, t string) string {
    minWin := ""
    minWinLen := math.MaxInt
    tCounts := make(map[byte]int)
    wCounts := make(map[byte]int)

    for i := range t {
        tCounts[t[i]]++
    }

    targetIncluded := func() bool {
        for k, v := range tCounts {
            if wCounts[k] < v {
                return false
            }
        }
        return true
    }

    for l, r := 0, 0; r < len(s); r++ {
        wCounts[s[r]]++
        for targetIncluded() {
            if minWinLen > r - l + 1 {
                minWinLen = r - l + 1
                minWin = s[l:r+1]
            }
            wCounts[s[l]]--
            l++
        }
    }
    return minWin
}
