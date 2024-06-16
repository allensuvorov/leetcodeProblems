func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    if s == goal && hasDuplicate(s) {
        return true
    }

    diff := []int{}
    for i := range s {
        if s[i] != goal[i] {
           diff = append(diff, i)
        }
    }
    return len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
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
