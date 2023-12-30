func countAsterisks(s string) int {
    count := 0
    betweenPair := false
    for i := range s {
        if s[i] == '|' {
            betweenPair = !betweenPair
        }
        if !betweenPair && s[i] == '*' {
            count++
        }
    }
    return count
}
