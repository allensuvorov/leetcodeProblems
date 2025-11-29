type StreamChecker struct {
    stream []byte
    t trie
}

type trieNode struct {
    children [26]*trieNode
    isEndOfWord bool
}

type trie struct {
    root *trieNode
}

func newTrie () trie {
    t := trie {
        root: new(trieNode),
    }
    return t
}

func (t *trie) insert(word string) {
    now := t.root
    for i := len(word)-1; i >= 0; i-- {
        c := word[i]
        if now.children[c - 'a'] == nil {
            now.children[c - 'a'] = new(trieNode)
        }
        now = now.children[c - 'a']
    }
    now.isEndOfWord = true
}


func (t *trie) query(stream []byte) bool {
    now := t.root
    for i := len(stream) - 1; i >= 0; i-- {
        c := stream[i]
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

func Constructor(words []string) StreamChecker {
    var stream []byte
    t := newTrie()
    for _, word := range words {
        t.insert(word)
    }
    return StreamChecker{stream, t}
}


func (this *StreamChecker) Query(letter byte) bool {
    this.stream = append(this.stream, letter)
    // fmt.Println(string(this.stream))
    return this.t.query(this.stream)
}


/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
