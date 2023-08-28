func findSubstring(s string, words []string) []int {
    n := len(s)
    wordLenth := len(words[0])
    m := len(words) * wordLenth
    res := []int{}
    if m > n {
        return res
    }

    target := map[string]int{}
    window := map[string]int{}
    for _, v := range words {
        target[v]++
    }
    
    have, need := 0, len(target)
    l := 0
    for r := 0; r < n - wordLenth; r += wordLenth {
        word := s[r : r + wordLenth]
        if _, ok := target[word]; ok {
            window[word]++
            fmt.Println(window, have, need)

            if window[word] == target[word] {
                have++
            }
            if window[word] > target[word] {
                l = r
                have = 0
                window = map[string]int{}
                window[word]++
                if window[word] == target[word] {
                    have++
                }
            }
        } else {
            have = 0
            window = map[string]int{}
            l = r + wordLenth
        }

        if have == need {
            res = append(res, l)
            //fmt.Println(window, have, need)
            word := s[l:l + wordLenth]
            window[word]--
            have--
            l += wordLenth 
        }
    }
    return res
}
