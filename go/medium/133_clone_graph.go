/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    clones := make(map[*Node]*Node) 

    var dfs func(node *Node) *Node
    dfs = func(node *Node) *Node {
        if clone, ok := clones[node]; ok {
            return clone
        }
        clone := &Node{node.Val, []*Node{}}
        clones[node] = clone
        for _, nei := range node.Neighbors {
            clone.Neighbors = append(clone.Neighbors, dfs(nei))
            fmt.Println(clone)
        }
        return clone
    }
    
    return dfs(node)
}   
