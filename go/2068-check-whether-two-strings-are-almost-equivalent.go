func checkAlmostEquivalent(word1 string, word2 string) bool {
    // ---0+++
    diffMap := make(map[byte]int)
    for i := range word1 {
        diffMap[word1[i]]++
        diffMap[word2[i]]--
    }
    for i := range diffMap {
        if diffMap[i] > 3 || diffMap[i] < -3 {
            return false
        }
    }
    return true
}
