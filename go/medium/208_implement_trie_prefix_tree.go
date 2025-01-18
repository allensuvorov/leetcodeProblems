type Trie struct {
    root *node
}

type node struct {
    children [26]*node
    isWord bool
}


func Constructor() Trie {
    return Trie{&node{}}
}


func (this *Trie) Insert(word string) {
    now := this.root
    for _, c := range word {
        if now.children[c - 'a'] == nil {
            now.children[c - 'a'] = new(node)
        }
        now = now.children[c - 'a']
    }
    now.isWord = true
}


func (this *Trie) Search(word string) bool {
    now := this.root
    for _, c := range word {
        if now.children[c - 'a'] == nil {
            return false
        }
        now = now.children[c - 'a']
    }
    return now.isWord
}


func (this *Trie) StartsWith(prefix string) bool {
    now := this.root
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
