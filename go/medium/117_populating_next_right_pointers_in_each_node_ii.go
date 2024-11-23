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
	if root == nil {
        return nil
    }
    tail, head := root, root
    levelSize := 1
    for head != nil {
        newSize := 0

        for i := range levelSize {
            if head.Left != nil {
                tail.Next = head.Left
                tail = tail.Next
                newSize++
            }
            if head.Right != nil {
                tail.Next = head.Right
                tail = tail.Next
                newSize++
            }
            next := head.Next
            if i == levelSize - 1 {
                head.Next = nil
            }
            head = next
        }
        levelSize = newSize
    }
    return root
}
