/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }
    
    if head.Next == nil {
        return head
    }

    dummy1, dummy2 := &ListNode{}, &ListNode{}
    cur1, cur2 := dummy1, dummy2
    
    for cur := head; cur != nil; cur = cur.Next {
        if cur.Val < x {
            cur1.Next = &ListNode{}
            cur1 = cur1.Next
            cur1.Val = cur.Val
        } else {
            cur2.Next = &ListNode{}
            cur2 = cur2.Next
            cur2.Val = cur.Val
        }
    }
    cur1.Next = dummy2.Next
    return dummy1.Next
}
