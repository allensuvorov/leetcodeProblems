func closeStrings(word1 string, word2 string) bool {
    charCount1 := [26]int{}
    for i := range word1 {
        charCount1[word1[i] - 'a']++
    }

    charCount2 := [26]int{}
    for i := range word2 {
        charCount2[word2[i] - 'a']++
    }

    for i := range 26 {
        if charCount1[i] == 0 && charCount2[i] != 0 {
            return false
        }
        if charCount1[i] != 0 && charCount2[i] == 0 {
            return false
        }
    }

    freqCount1 := map[int]int{}
    freqCount2 := map[int]int{}
    for i := range 26 {
        freqCount1[charCount1[i]]++
        freqCount2[charCount2[i]]++
    }

    for k, v := range freqCount1{
        if freqCount2[k] != v {
            return false
        }
    }
    return true
}
