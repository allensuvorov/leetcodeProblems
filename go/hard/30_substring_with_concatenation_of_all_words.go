func findSubstring(s string, words []string) []int {
    wordLenth := len(words[0])
    n := len(s)
    m := len(words)

    // is there a better solution then n * m ?
    res := []int{}
    if m * wordLenth > n {
        return res
    }
    
    target := map[string]int{}
    window := map[string]int{}
    for _, v := range words {
        target[v]++
    }
    
    have, need := 0, len(target)
    l, r := 0, 0

    nextLetter := func() {        
        have = 0
        window = map[string]int{}
        l++
        r = l
    }
    for l <= n - m * wordLenth {

        word := s[r : r + wordLenth]
        
        if _, ok := target[word]; ok {
            window[word]++
            if window[word] == target[word] {
                have++
            }
            r += wordLenth
            if window[word] > target[word] {
                nextLetter()
            }
        } else {
            nextLetter()
        }
        // fmt.Println(l, r, window, have, need)

        if have == need {
            res = append(res, l)
            nextLetter()
        }
    }
    return res
}
