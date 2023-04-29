func countVowelSubstrings(word string) int {
    res := 0
    for l := range word {
        currentVS := map[byte]bool{}
        for r := l; r < len(word) && isVowel(word[r]); r++ {
            currentVS[word[r]] = true
            if len(currentVS) == 5 {
                res++
            }
        }
    }
    return res
}

func isVowel (b byte) bool {
    return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
