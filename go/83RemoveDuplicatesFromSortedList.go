/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    ptr := head
    res := head
    for head != nil && ptr.Next != nil {
        if ptr.Next.Val != ptr.Val {
            res.Next = ptr.Next
            ptr = ptr.Next
            res = res.Next
        } else {
            ptr = ptr.Next
            res.Next = nil
        }
        
        fmt.Println(res)
    }
    return head
}
