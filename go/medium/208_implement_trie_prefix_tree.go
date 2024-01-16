type Trie struct {
    root *Node
}


type Node struct {
    children [26] *Node
    isEnd bool
}

func Constructor() Trie {
    return Trie{new(Node)}
}


func (this *Trie) Insert(word string)  {
    now := this.root
    for _, c := range word {
        i := c - 'a'
        if now.children[i] == nil {
            now.children[i] = new(Node)
        }
        now = now.children[i]
    }
    now.isEnd = true
}


func (this *Trie) Search(word string) bool {
    now := this.root
    for _, c := range word {
        i := c - 'a'
        if now.children[i] == nil {
            return false
        }
        now = now.children[i]
    }
    return now.isEnd 
}


func (this *Trie) StartsWith(prefix string) bool {
    now := this.root
    for _, c := range prefix {
        i := c - 'a'
        if now.children[i] == nil {
            return false
        }
        now = now.children[i]
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
