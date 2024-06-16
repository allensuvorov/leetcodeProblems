func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    if s == goal && hasDuplicate(s) {
        return true
    }

    firstDiff := 0
    diffCount := 0
    for i := range s {
        if s[i] != goal[i] {
            diffCount++
            if diffCount == 1 {
                firstDiff = i
            }
            if diffCount == 2 {
                if s[i] != goal[firstDiff] || s[firstDiff] != goal[i] {
                    return false
                }
            }
            if diffCount == 3 {
                return false
            }   
        }
    }
    return diffCount == 2
}

func hasDuplicate(s string) bool {
    charCount := map[rune]int{}
    for _, v := range s {
        if charCount[v] == 1 {
            return true
        }
        charCount[v]++
    }
    return false
}
