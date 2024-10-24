/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    evenHead := head.Next
    odd := head
    even := head.Next
    
    for odd.Next != nil && even.Next != nil {
        
        odd.Next = even.Next
        even.Next = even.Next.Next
        
        odd = odd.Next
        even = even.Next

    }
    odd.Next = evenHead
    return head
}
