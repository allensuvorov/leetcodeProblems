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
        if now.isWord {
            return true
        } else {
            return false
        }
    }

    if word[i] == '.' {
        i++
        for _, node := range now.children {
            if node != nil {
                if this.dfs(node, word, i) {
                    return true
                }
            }
        }
    } else {
        c := word[i]
        if now.children[c - 'a'] == nil {
            return false
        }
        i++
        now = now.children[c - 'a']
        // fmt.Println(i, string(word[i]))
        return this.dfs(now, word, i)
    }
    return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
