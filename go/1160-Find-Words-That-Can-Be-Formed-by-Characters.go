func countCharacters(words []string, chars string) int {
    cm := map[rune]int{}
    for _, v := range chars {
        cm[v]++
    }
    ans := 0
    for _, w := range words {
        // cm has c's for w
        isGood := true
        wcm := map[rune]int{}
        for _, c := range w {
            wcm[c]++
            if wcm[c] > cm[c] {
                isGood = false
                break
            }
        }
        if isGood {
            ans += len(w)
        }
    }
    return ans
}
