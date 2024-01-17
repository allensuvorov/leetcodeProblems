func suggestedProducts(products []string, searchWord string) [][]string {
    t := newTrie()
    res := [][]string{}

    for _, p := range products {
        t.Insert(p)
    }

    pref := []rune{}
    for _, c := range searchWord {
        pref = append(pref, c)
        res = append(res, t.getWordsStartingWith(pref))
    }
    return res
}


type Trie struct {
    Root *Node
}

    
type Node struct {
    Children [26] *Node // array for edges - to sort lexicographically
    IsWord bool
}


func newTrie() Trie {
    return Trie{new(Node)}
}


func (this *Trie) Insert(product string) {
    now := this.Root
    for _, c := range product {
        i := c - 'a'
        if now.Children[i] == nil {
            now.Children[i] = new(Node)
        }
        now = now.Children[i]
    }
    now.IsWord = true
}

func (this *Trie) getWordsStartingWith(pref []rune) []string{
    now := this.Root
    res := []string{}
    // Check: are there any words with that prefix?
    for _, c := range pref {
        if now.Children[c - 'a'] == nil { 
            return res
        }
        now = now.Children[c - 'a']
    }
    this.dfsWithPrefix(now, pref, &res)
    return res
}

func (this *Trie) dfsWithPrefix(now *Node, pref []rune, result *[]string) {
    if len(*result) == 3 {
        return
    }
    if now.IsWord {
        *result = append(*result, string(pref))
    }
    for i, node := range now.Children {
        if node != nil {
            this.dfsWithPrefix(node, append(pref, rune(i + 'a')), result)
            
        }
    }
}
