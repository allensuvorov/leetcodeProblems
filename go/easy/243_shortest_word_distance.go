func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    minDist := len(wordsDict)
    lastWord1Pos := -1 
    lastWord2Pos := -1

    for i, v := range wordsDict {
        if v == word1 {
            lastWord1Pos = i
            if lastWord2Pos >= 0 {
                minDist = min(minDist, lastWord1Pos - lastWord2Pos)
            }
        }
        if v == word2 {
            lastWord2Pos = i
            if lastWord1Pos >= 0 {
                minDist = min(minDist, lastWord2Pos - lastWord1Pos)
            }
        }
    }
    
    return minDist
}
