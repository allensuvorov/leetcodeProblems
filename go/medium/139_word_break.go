type Trie struct {
    Root *TrieNode
}

type TrieNode struct {
    Children [26]*TrieNode
    IsWord bool
}

func wordBreak(s string, wordDict []string) bool {

    // trie - prefix tree
    t := Trie{&TrieNode{}}

    for _, word := range wordDict {
        now := t.Root
        for _, v := range word {
            if now.Children[v - 'a'] == nil {
                now.Children[v - 'a'] = &TrieNode{}
            }
            now = now.Children[v - 'a']
        }
        now.IsWord = true
    }
    
    // backtracking
    var backtrack func(node *TrieNode, s string) bool
    backtrack = func(node *TrieNode, s string) bool {
        if len(s) == 0 {
            return true
        }
        for i, v := range s{
            if node.Children[v - 'a'] == nil {
                return false
            }
            node = node.Children[v - 'a']
            if node.IsWord && backtrack(t.Root, s[i + 1:]) {
                return true
            }
        }
        return false
    }
    return backtrack(t.Root, s)

}
