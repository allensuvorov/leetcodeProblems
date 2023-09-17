/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    ptrInd := map[*Node]int{}
    indPtr := []*Node{}
    dummy := &Node{}
    l2 := dummy
    l1 := head
    for i := 0; l1 != nil; i++ {
        l2.Next = &Node{}
        l2 = l2.Next
        l2.Val = l1.Val
        
        ptrInd[l1] = i
        indPtr = append(indPtr, l2)
        l1 = l1.Next
    }
    l1 = head
    l2 = dummy.Next
    for l1 != nil {
        if l1.Random != nil {
            l2.Random = indPtr[ptrInd[l1.Random]]
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return dummy.Next
}
