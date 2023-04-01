/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    
    res := head
    
    if head != nil {
        for ptr := head; ptr.Next != nil; ptr = ptr.Next {
            if ptr.Next.Val != ptr.Val {
                res.Next = ptr.Next
                res = res.Next
            }
        }
        res.Next = nil
    }
    return head
}
