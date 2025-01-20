type trie struct {
    children [26]*trie
    isWord bool
}

func (t *trie) insert(word string) {
    now := t
    for _, c := range word {
        if now.children[c - 'a'] == nil {
            now.children[c - 'a'] = new(trie)
        }
        now = now.children[c - 'a']
    }
    now.isWord = true
}

func (t *trie) searchThree(prefix string) []string {
    res := make([]string, 0)
    word := []byte(prefix)

    var dfs func(t *trie) 
    dfs = func(t *trie) {
        if len(res) == 3 {
            return
        }
        
        if t.isWord {
            res = append(res, string(word))
        }
        
        for i, node := range t.children {
            if node != nil {
                word = append(word, byte(i + 'a'))
                dfs(node)
                word = word[:len(word)-1]
            } 
        }
    }
    dfs(t)
    return res
}

func suggestedProducts(products []string, searchWord string) [][]string {
    t := new(trie)
    for _, v := range products {
        t.insert(v)
    }

    res := make([][]string, len(searchWord))
    
    now := t
    for i := 0; i < len(searchWord) && now.children[searchWord[i] - 'a'] != nil; i++ {
        now = now.children[searchWord[i] - 'a']
        res[i] = now.searchThree(searchWord[:i+1])
    }

    return res
}
