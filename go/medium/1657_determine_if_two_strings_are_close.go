func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    var charSet1, charSet2 [26]bool
    for i := range word1 {
        charSet1[word1[i] - 'a'] = true
        charSet2[word2[i] - 'a'] = true
    }

    if charSet1 != charSet2 {
        return false
    }

    var charCount1, charCount2 [26]int

    for i := range word1 {
        charCount1[word1[i] - 'a']++
        charCount2[word2[i] - 'a']++
    }

    freqCount1, freqCount2 := map[int]int{}, map[int]int{}

    for i := range charCount1 {
        freqCount1[charCount1[i]]++
        freqCount2[charCount2[i]]++
    }

    for freq1, count1 := range freqCount1 {
        if count2, ok := freqCount2[freq1]; !ok || count1 != count2 {
            return false
        }
    }

    return  true
}
