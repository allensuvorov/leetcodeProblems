/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeElements(head *ListNode, val int) *ListNode {
    
    for head!=nil {
        if head.Val == val{
            head = head.Next
        } else {
            break
        }
    }
    
    pre := head
    
    for ptr := head.Next ; ptr != nil; ptr = ptr.Next {
        if ptr.Val == val {
            pre.Next = ptr.Next
        }
        pre = ptr
    }
    return head
}