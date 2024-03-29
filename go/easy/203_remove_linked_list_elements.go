/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{ Next: head }
    prev := dummy
    for prev.Next != nil {
        if prev.Next.Val == val {
            prev.Next = prev.Next.Next
        } else {
            prev = prev.Next
        }
    }
    return dummy.Next
}
