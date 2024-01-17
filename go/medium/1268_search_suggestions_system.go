func suggestedProducts(products []string, searchWord string) [][]string {
    t := newTrie()
    for _, p := range products {
        t.Insert(p)
    }
    return t.FindProducts(searchWord)
}


type Trie struct {
    Root *Node
}

    
type Node struct {
    Children [26] *Node // array for edges - to sort lexicographically
    IsEnd bool
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
    now.IsEnd = true
}


func (this *Trie) FindProducts(searchWord string) [][]string{
    res := [][]string{}
    now := this.Root
    pref := []rune{}
    for _, c := range searchWord {
        i := c - 'a'
        pref = append(pref, c)
        count := 0
        // fmt.Printf("%v \n",string(pref))
        if now.Children[i] != nil {
            // fmt.Println(now.Children[i].Children)
            threeProducts := make([]string, 3)
            this.FindThreeProducts(now.Children[i], pref, threeProducts, &count)
            res = append(res, threeProducts[:count])
            now = now.Children[i]
        } else {
            res = append(res, []string{})
        }
    }
    return res
}

func (this *Trie) FindThreeProducts(now *Node, pref []rune, threeProducts []string, count *int) {
    
    if now.IsEnd {
        // fmt.Println(string(pref), *count)
        threeProducts[*count] = string(pref)
        *count++
    }
    for i, node := range now.Children {
        if node != nil {
            pref := append(pref, rune(i + 'a'))
            // fmt.Println(string(rune(i + 'a')), string(pref), *count)
            now = node
            this.FindThreeProducts(now, pref, threeProducts, count)
        }
        if *count == 3 {
            return
        }
    }
}
