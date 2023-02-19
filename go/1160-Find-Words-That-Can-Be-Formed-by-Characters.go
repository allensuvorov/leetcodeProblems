func countCharacters(words []string, chars string) int {
    cm := map[byte]int{}
    for i := range chars {
        cm[chars[i]]++
    }

    ans := 0
    for i := range words {
        isGood := true
        wcm := map[byte]int{}

        for j := range words[i] {
            wcm[words[i][j]]++
            if wcm[words[i][j]] > cm[words[i][j]] {
                isGood = false
                break
            }
        }
        if isGood {
            ans += len(words[i])
        }
    }
    return ans
}
