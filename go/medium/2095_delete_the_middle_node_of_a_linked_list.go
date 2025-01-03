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
    }
    prev.Next = slow.Next
    return head
}

// redo
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

    listLen := 0
    for runner := head; runner != nil; runner = runner.Next {
        listLen++
    }

    half := listLen / 2

    runner := head
    for range half - 1 {
        runner = runner.Next
    }


    runner.Next = runner.Next.Next 
        
    return head
}
