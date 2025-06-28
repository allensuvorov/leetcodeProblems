func calculateScore(s string) int64 {
    indices := make([][]int, 26)
    var totalScore int64
    for i, c := range s {
        mirror := 'z'- 'a' - (c - 'a')
        if l := len(indices[mirror]); l > 0 {
            closestUnmarkedIndex := indices[mirror][l-1]
            indices[mirror] = indices[mirror][:l-1] // pop
            score := i - closestUnmarkedIndex
            totalScore += int64(score)
        } else {
            indices[c - 'a'] = append(indices[c - 'a'], i)
        }
    }
    return totalScore
}
