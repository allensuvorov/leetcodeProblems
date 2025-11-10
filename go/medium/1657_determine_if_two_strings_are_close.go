func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    var charCount1, charCount2 [26]int

    for i := range word1 {
        charCount1[word1[i] - 'a']++
        charCount2[word2[i] - 'a']++
    }

    for i := range 26 {
        if charCount1[i] * charCount2[i] == 0 && charCount1[i] != charCount2[i]{
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
