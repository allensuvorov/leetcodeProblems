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
    list1 := root
    for list1 != nil {
        dummy := &Node{}
        list2 := dummy
        for list1 != nil {
            if list1.Left != nil {
                list2.Next = list1.Left
                list2 = list2.Next
            }
            if list1.Right != nil {
                list2.Next = list1.Right
                list2 = list2.Next
            }
            list1 = list1.Next
        }
        list1 = dummy.Next
    }
    return root
}
