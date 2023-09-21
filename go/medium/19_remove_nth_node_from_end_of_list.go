/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        return nil
    }
    ptrs := [33]*ListNode{}
    cur := head
    // fill the list of ptrs
    i := 0
    for i = 1; i < 31 && cur != nil; i++ {
        ptrs[i] = cur
        cur = cur.Next
    }
    i--
    if i - n == 0 {
        return head.Next
    }
    // go directly by pointers and reconnect nodes
    ptrs[i-n].Next = ptrs[i-n+2]
    return head
}
