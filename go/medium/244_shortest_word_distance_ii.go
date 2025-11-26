type WordDistance struct {
    wordIndexes map[string][]int
}


func Constructor(wordsDict []string) WordDistance {
    wordIndexes := map[string][]int{}
    for i, v := range wordsDict {
        if _, ok := wordIndexes[v]; !ok {
            wordIndexes[v] = []int{}
        }
        wordIndexes[v] = append(wordIndexes[v], i)
    }
    return WordDistance{wordIndexes}
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    inds1 := this.wordIndexes[word1]
    inds2 := this.wordIndexes[word2]
    // no point search forward from the bigger one - it's only going to get bigger
    minDist := math.MaxInt
    for i, j := 0, 0; i < len(inds1) && j < len(inds2); {
        minDist = min(minDist, abs(inds1[i] - inds2[j]))
        if inds1[i] < inds2[j] {
            i++
        } else {
            j++
        }
    }
    return minDist
}


func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */

