type WordDictionary struct {
    root *Node
}

type Node struct {
    isWord bool
    children [26]*Node
}


func Constructor() WordDictionary {
    return WordDictionary{new(Node)}
}


func (this *WordDictionary) AddWord(word string)  {
    now := this.root
    for _, c := range word {
        i := c - 'a'
        if now.children[i] == nil {
            now.children[i] = new(Node)
        }
        now = now.children[i]
    }
    now.isWord = true
}


func (this *WordDictionary) Search(word string) bool {
    return this.dfs(this.root, word, 0)
}

func (this *WordDictionary) dfs(now *Node, word string, i int) bool {
    if i == len(word) {
        return now.isWord
    }

    if word[i] == '.' {
        for _, node := range now.children {
            if node != nil {
                if this.dfs(node, word, i + 1) {
                    return true
                }
            }
        }
    } else {
        if node := now.children[word[i] - 'a']; node == nil {
            return false
        } else {
            return this.dfs(node, word, i + 1)
        }
    }
    return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
