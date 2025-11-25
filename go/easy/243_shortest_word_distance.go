func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    minDist := len(wordsDict)
    mostRecentPosWord1 := -1 
    mostRecentPosWord2 := -1

    for i, v := range wordsDict {
        if v == word1 {
            mostRecentPosWord1 = i
            if mostRecentPosWord2 >= 0 {
                minDist = min(minDist, mostRecentPosWord1 - mostRecentPosWord2)
            }
        }
        if v == word2 {
            mostRecentPosWord2 = i
            if mostRecentPosWord1 >= 0 {
                minDist = min(minDist, mostRecentPosWord2 - mostRecentPosWord1)
            }
        }
    }
    
    return minDist
}
