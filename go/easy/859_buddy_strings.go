func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    prevIndex := 0
    difCount := 0
    for i := range s {
        if s[i] != goal[i] {
            difCount++
            if difCount == 1 {
                prevIndex = i
            }
            if difCount == 2 {
                if s[prevIndex] != goal[i] || s[i] != goal[prevIndex] {
                    return false
                }
            }
            if difCount == 3 {
                return false
            }
        }
    }
    if difCount < 2 {
        return false 
    }
    return true
}
