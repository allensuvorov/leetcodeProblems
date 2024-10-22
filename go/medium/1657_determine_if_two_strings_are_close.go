func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    if getSet(word1) != getSet(word2) {
        return false
    }

    if getFreqCount(word1) != getFreqCount(word2) {
        return false
    }
    return true
}

func getSet(s string) [26]bool {
    byteSet := [26]bool{}
    for _, v := range s {
        byteSet[v - 'a'] = true
    }
    return byteSet
}

func getFreqCount (s string) [1e5+1]int {
    byteCount := [26]int{}
    freqCount := [1e5+1]int{}
    for _, v := range s {
        byteCount[v - 'a']++
    }

    for _, v := range byteCount {
        freqCount[v]++
    }
    return freqCount
}
