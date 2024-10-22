func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    if getSet(word1) != getSet(word2) {
        return false
    }

    fcfc1 := getFreqCount(word1) 
    fcfc2 := getFreqCount(word2)

    for k, v := range fcfc1 {
        if fcfc2[k] != v {
            return false
        }
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

func getFreqCount (s string) map[int]int {
    byteCount := [26]int{}
    for _, v := range s {
        byteCount[v - 'a']++
    }

    freqCount := map[int]int{}
    for _, v := range byteCount {
        freqCount[v]++
    }
    return freqCount
}
