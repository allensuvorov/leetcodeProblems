type Trie struct {
    children [26]*Trie
    isWord bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string) {
    now := this
    for _, c := range word {
        if now.children[c - 'a'] == nil {
            now.children[c - 'a'] = new(Trie)
        }
        now = now.children[c - 'a']
    }
    now.isWord = true
}


func (this *Trie) Search(word string) bool {
    now := this
    for _, c := range word {
        if now.children[c - 'a'] == nil {
            return false
        }
        now = now.children[c - 'a']
    }
    return now.isWord
}


func (this *Trie) StartsWith(prefix string) bool {
    now := this
    for _, c := range prefix {
        if now.children[c - 'a'] == nil {
            return false
        }
        now = now.children[c - 'a']
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
