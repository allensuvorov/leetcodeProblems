func findWords(board [][]byte, words []string) []string {
    t := Trie{new(Node)}
    
    rows, cols := len(board), len(board[0])

    var dfs func(int, int, int, *Node)
    dfs = func(r, c, i int, now *Node) {
        if i == 10 {
            return
        }
        if r < 0 || r == rows || c < 0 || c == cols || board[r][c] == '1' {
            return
        }
        char := board[r][c]

        // trie insert
        if now.children[char - 'a'] == nil {
            now.children[char - 'a'] = new(Node)
        }
        now = now.children[char - 'a']

        // matrix traverse
        board[r][c] = '1'
        dfs(r + 1, c, i + 1, now)
        dfs(r - 1, c, i + 1, now)
        dfs(r, c + 1, i + 1, now)
        dfs(r, c - 1, i + 1, now)
        board[r][c] = char
    }

    for r := range board {
        for c := range board[0] {
            dfs(r, c, 0, t.root)
        }
    }

    res := []string{}
    for _, w := range words {
        if t.Search(w) {
            res = append(res, w)
        }
    }
    return res
}

type Trie struct {
    root *Node
}

type Node struct {
    children [26]*Node
}

func (t *Trie) Search(word string) bool {
    now := t.root
    for _, c := range word {
        if now.children[c - 'a'] == nil {
            return false
        }
        now = now.children[c - 'a']
    }
    return true
}
