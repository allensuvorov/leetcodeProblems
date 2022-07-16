/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    a1, a2 := headA, headA
    b1, b2 := headB, headB
    
    for a1 != nil && b1 != nil{
        if a1.Next != nil {
            a1 = a1.Next
        } else {
            b2 = b2.Next 
        }
        
        if b2.Next != nil{
            b2 = b2.Next    
        } elst {
            a2 = n2.Next
        }
    }
    if a1 != b1 {
        return nil
    }
    
    return
}