/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }
    fast := head
    slow := head
    prev := head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        if fast.Next != nil {
            fast = fast.Next.Next
        }
        fmt.Println(prev.Val, slow.Val)
    }
    prev.Next = slow.Next
    return head
}
