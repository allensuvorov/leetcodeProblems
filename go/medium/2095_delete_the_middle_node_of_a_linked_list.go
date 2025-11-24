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

    slow, fast := head, head.Next.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    slow.Next = slow.Next.Next // delete middle node
    return head
}

// with prev pointer
func deleteMiddle(head *ListNode) *ListNode {
    // special case - one node list
    if head.Next == nil {
        return nil
    }

    // find middle node and prev node (find prev and next)
    var prev *ListNode
    slow := head // middle one
    fast := head // finds the end of list at double speed

    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    } 

    // remove middle node
    prev.Next = slow.Next // link
    return head
}
