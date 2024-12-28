func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    charCount1, charCount2 := make(map[byte]int), make(map[byte]int)

    for i := range word1 {
        charCount1[word1[i]]++
        charCount2[word2[i]]++
    }

    if len(charCount1) != len(charCount2) {
        return false
    }
    
    // same chars
    for k := range charCount1 {
        if _, exists := charCount2[k]; !exists {
            return false
        }
    }

    // same counts
    countCount1 := make(map[int]int)
    for _, v := range charCount1 {
        countCount1[v]++
    }

    countCount2 := make(map[int]int) 
    for _, v := range charCount2 {
        countCount2[v]++
    }

    for k, v1 := range countCount1 {
        if v2, ok := countCount2[k]; !ok || v1 != v2 {
            return false
        }
    }
    return true
}
