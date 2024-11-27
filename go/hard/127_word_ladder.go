func ladderLength(beginWord string, endWord string, wordList []string) int {
    steps := 0
    q := []string{beginWord}
    visited := map[string]bool{beginWord:true}
    
    wordSet := make(map[string]bool, len(wordList))
    for _, word := range wordList {
        wordSet[word] = true
    }

    for len(q) > 0 {
        size := len(q)
        steps++
        for range size {
            now := q[0]
            q = q[1:]
            if now == endWord {
                return steps
            }
            for i := range now {
                for char := byte('a'); char <= 'z'; char++ {
                    if char != now[i] {
                        nextWord := now[:i] + string(char) + now[i+1:]
                        if _, ok := wordSet[nextWord]; ok {
                            if !visited[nextWord] {
                                visited[nextWord] = true
                                q = append(q, nextWord)
                            }
                        }
                    }
                }
            }
        }
    }
    return 0
}


