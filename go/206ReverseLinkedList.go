/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    
    // save pointers
    var ptr1 *LIstNode = head.Next
    var ptr2 *ListNode
    
    
    head.Next = ptr2
    
    return head
}
