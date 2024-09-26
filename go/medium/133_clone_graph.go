/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }
    
    clone := new(Node)
    clone.Val = node.Val
    
    q := []*Node{node}
    cloneQ := []*Node{clone}

    visited := make(map[*Node]bool)

    for len(q) > 0 {
        neiSize := len(q)
        for range neiSize {
            now := q[0]
            visited[now] = true
            fmt.Println(now.Val, "is now visited")
            q = q[1:]

            cloneNow := cloneQ[0]
            cloneQ = cloneQ[1:]

            for _, nei := range now.Neighbors {
                if !visited[nei] {
                    fmt.Println("working on nei:", nei.Val)
                    q = append(q, nei)

                    cloneNei := new(Node)
                    cloneQ = append(q, cloneNei)
                    cloneNei.Val = nei.Val

                    cloneNow.Neighbors = append(cloneNow.Neighbors, cloneNei)
                    cloneNei.Neighbors = append(cloneNei.Neighbors, cloneNow)
                    if cloneNow.Val == 2 {
                        fmt.Println("Clone 2 neighbros", cloneNow.Neighbors)
                    }
                }
            }
            
            fmt.Print(cloneNow.Val, " - ")
            for _, nei := range cloneNow.Neighbors {
                fmt.Print(nei.Val, ", ")
            }
            fmt.Println()

        }
    }
    return clone
}   
