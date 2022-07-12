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
    
    var rev, next, prev *ListNode
    
    rev = head
    next = rev.Next // (2) next address
    rev.Next = nil // (1)->null
    prev = rev // (1) current address
    
    for rev = next; rev.Next != nil; rev = next{
        // rev = 2
        next = rev.Next // address of (3)
        rev.Next = prev // (2)->(1) pointer
        prev = rev // save (2) address
        
        fmt.Println("end loop", rev.Val, rev.Next.Val)
    }
    rev.Next = prev
    
    return rev
}
