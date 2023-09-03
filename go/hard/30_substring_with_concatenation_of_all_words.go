func findSubstring(s string, words []string) []int {
    if len(words) == 0 {
        return nil
    }
    
    res := []int{}
    wordLen := len(words[0])
    numWords := len(words)
    wordsMap := map[string]int{}
    for _, w := range words {
        wordsMap[w]++
    }

    for i := 0; i < wordLen; i++ {
        start := i
        end := i
        counter := 0
        tmpMap := map[string]int{}

        for end+wordLen <= len(s) {
            word := s[end : end+wordLen]
            end += wordLen

            if wordsMap[word] != 0 {
                tmpMap[word]++

                if tmpMap[word] <= wordsMap[word] {
                    counter++
                }

                if counter == numWords {
                    res = append(res, start)
                }

                if end-start == wordLen*numWords {
                    if tmpMap[s[start:start+wordLen]] <= wordsMap[s[start:start+wordLen]] {
                        counter--
                    }
                    tmpMap[s[start:start+wordLen]]--
                    start += wordLen
                }
            } else {
                tmpMap = map[string]int{}
                start = end
                counter = 0
            }
        }
    }

    return res
}
