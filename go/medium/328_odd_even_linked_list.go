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
    evenHead = head.Next
    odd := &ListNode{}
    even := &ListNode{}
    
    isOdd := true
    for runner := head; runner != nil; runner = runner.Next {
        if isOdd {
            odd.Next = runner
            odd = odd.Next
        } else {
            even.Next = runner
            even = even.Next
        }
        isOdd = !isOdd
    }
    odd.Next = evenHead
    even.Next = nil
    return head
}
