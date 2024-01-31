/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    prev := head
    head1 := head
    for cur := head.Next; cur != nil; cur = cur.Next {
        if cur.Val < prev.Val {
            prev1 := &ListNode{}
            cur1 := head1
            for cur.Val > cur1.Val {
                prev1 = cur1
                cur1 = cur1.Next
            }
            prev1.Next = cur
            prev.Next = cur.Next
            cur.Next = cur1
            if cur1 == head1 {
                head1 = cur
            }
        }
        prev = cur 
    }
    return head1
}
