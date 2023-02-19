func countCharacters(words []string, chars string) int {
    cm := map[rune]int{}
    for _, v := range chars {
        cm[v]++
    }
    ans := 0
    for _, w := range words {
        l := len(w)
        for _, c := range w {
            if strings.Count(w, string(c)) > cm[c] {
                l = 0
                break
            }
        }
        ans += l
    }
    return ans
}
