/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    q1 := root // currentQueue
    for q1 != nil {
        dummy := &Node{} // nextQueue
        q2 := dummy
        for q1 != nil {
            if q1.Left != nil {
                q2.Next = q1.Left
                q2 = q2.Next
            }
            if q1.Right != nil {
                q2.Next = q1.Right
                q2 = q2.Next
            }
            q1 = q1.Next
        }
        q1 = dummy.Next
    }
    return root
}
