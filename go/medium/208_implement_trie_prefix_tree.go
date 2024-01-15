type Trie struct {
    root *TrieNode
}


type TrieNode struct {
    children map[rune] *TrieNode
    isEnd bool
}


func NewTrieNode() TrieNode {
    nm := make(map[rune] *TrieNode)
    return TrieNode {children: nm}
}


func Constructor() Trie {
    tn := NewTrieNode()
    return Trie {&tn}
}


func (this *Trie) Insert(word string)  {
    now := this.root
    for _, c := range word {
        if _, ok := now.children[c]; !ok {
            tn := NewTrieNode()
            now.children[c] = &tn
        }
        now = now.children[c]
    }
    now.isEnd = true
}


func (this *Trie) Search(word string) bool {
    now := this.root
    for _, c := range word {
        if _, ok := now.children[c]; !ok {
            return false
        }
        now = now.children[c]
    }
    return now.isEnd 
}


func (this *Trie) StartsWith(prefix string) bool {
    now := this.root
    for _, c := range prefix {
        if _, ok := now.children[c]; !ok {
            return false
        }
        now = now.children[c]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
