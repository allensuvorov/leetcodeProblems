func rotateString(s string, goal string) bool {
    return strings.Contains(s + s, goal) && len(s) == len (goal)
}

// someone's

func rotateString(s string, goal string) bool {

    for i := 0; i < len(s); i++ {
        rotated := s[i:] + s[:i]
        if rotated == goal {
            return true
        }
    }

    return false
}
