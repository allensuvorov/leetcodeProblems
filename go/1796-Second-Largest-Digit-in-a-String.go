func secondHighest(s string) int {
    var firstMax byte
    var secondMax byte
    exists := false
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            if firstMax == 0 {
                if firstMax < s[i] {
                    firstMax = s[i]
                }
            } else if firstMax < s[i] {
                secondMax = firstMax
                firstMax = s[i]
                exists = true
            } else if secondMax < s[i] && firstMax != s[i] {
                secondMax = s[i]
                exists = true
            }
        }
    }
    if exists {
        return int(secondMax-'0')
    } else {
        return -1
    }
}
