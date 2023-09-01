func findSubstring(s string, words []string) []int {
    wordLenth := len(words[0])
    n := len(s)
    m := len(words)*wordLenth

    // is there a better solution then n * m ?
    res := []int{}
    if m > n {
        return res
    }
    
    if sameLetter(s, words) {
        for i := 0; i < n-m + 1; i++ {
            res = append(res, i)
        }
        return res
    }

    target := map[string]int{}
    window := map[string]int{}
    for _, v := range words {
        target[v]++
    }
    
    have, need := 0, len(target)
    l, r := 0, 0
    for r <= n - wordLenth {
        //substring := s[r : r + m]

        word := s[r : r + wordLenth]
        
        if _, ok := target[word]; ok {
            window[word]++
            if window[word] == target[word] {
                have++
            }
            r += wordLenth

            if window[word] > target[word] { // what to do if one item overflows
                // start over from after the first entry of overflow letter
                have = 0
                window = map[string]int{}
                
                w := s[l : l + wordLenth]
                for w != word {
                    l += wordLenth
                    w = s[l : l + wordLenth]
                }
                l += wordLenth
                r = l
            }

        } else {
            have = 0
            window = map[string]int{}
            r++
            l = r
        }
        fmt.Println(l, r, window, have, need)

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

func sameLetter (s string, words []string) bool {
    for i := range s {
        if s[0] != s[i] {
            return false
        }
    }
    for _, w := range words {
        for i := range w {
            if words[0][0] != w[i] {
                return false
            }
        }
    }
    return true
}
