type StreamChecker struct {
    stream []byte
    root *trieNode
}


type trieNode struct {
    children [26]*trieNode
    isEndOfWord bool
}


func Constructor(words []string) StreamChecker {
    var stream []byte
    root := new(trieNode)
    for _, word := range words {
        now := root
        for i := len(word)-1; i >= 0; i-- {
            c := word[i]
            if now.children[c - 'a'] == nil {
                now.children[c - 'a'] = new(trieNode)
            }
            now = now.children[c - 'a']
        }
        now.isEndOfWord = true
    }
    return StreamChecker{stream, root}
}


func (this *StreamChecker) Query(letter byte) bool {
    this.stream = append(this.stream, letter)
    now := this.root
    for i := len(this.stream) - 1; i >= 0; i-- {
        c := this.stream[i]
        if now.children[c - 'a'] == nil {
            return false
        }
        now = now.children[c - 'a']
        if now.isEndOfWord {
            return true
        }
    }
    return false
}


/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
