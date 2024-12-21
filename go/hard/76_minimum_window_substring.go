func minWindow(s string, t string) string {
    minWin := ""
    minWinLen := math.MaxInt
    tCounts := make(map[byte]int)
    wCounts := make(map[byte]int)

    for i := range t {
        tCounts[t[i]]++
    }

    have, need := 0, len(t)

    for l, r := 0, 0; r < len(s); r++ {
        if tCounts[s[r]] > wCounts[s[r]] {
            have++
        }

        wCounts[s[r]]++

        for have == need {
            if minWinLen > r - l + 1 {
                minWinLen = r - l + 1
                minWin = s[l:r+1]
            }

            if tCounts[s[l]] == wCounts[s[l]] {
                have--
            }

            wCounts[s[l]]--
            l++
        }
    }
    return minWin
}
