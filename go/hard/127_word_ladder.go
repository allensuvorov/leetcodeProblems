func ladderLength(beginWord string, endWord string, wordList []string) int {
    // BFS till find endWord
    steps := 0
    q := []string{beginWord}
    visited := map[string]bool{beginWord:true}
    neighbors := adjTable(append(wordList, beginWord))
    for len(q) > 0 {
        size := len(q)
        steps++
        for range size {
            now := q[0]
            q = q[1:]
            if now == endWord {
                return steps
            }
            for _, nei := range neighbors[now] {
                if !visited[nei] {
                    visited[nei] = true
                    q = append(q, nei)
                }
            }
        }
    }
    return 0
}


func adjTable(wordList []string) map[string][]string{
    // generating all possible substituting letters
    // hit : *it, h*t, hi*
    neighbors := make(map[string][]string, len(wordList))
    for _, word := range wordList {
        neighbors[word] = make([]string,0)
    }

    for _, word := range wordList {
        for i := range word {
            for j := range 26 {
                char := byte(j + 'a')
                if char != word[i] {
                    wordBytes := []byte(word)
                    wordBytes[i] = char
                    genNei := string(wordBytes)
                    if _, ok := neighbors[genNei]; ok {
                        neighbors[genNei] = append(neighbors[genNei], word)
                    }
                }
            }
        }
    }

    return neighbors
}
