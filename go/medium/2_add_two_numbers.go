/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int
    l3 := &ListNode{}
    n1, n2, n3 := l1, l2, l3
    for {
        sum := carry
        if n1 != nil {
            sum += n1.Val
            n1 = n1.Next
        }
        if n2 != nil {
            sum += n2.Val
            n2 = n2.Next
        }
        n3.Val = sum % 10
        carry = sum / 10
        
        if n1 != nil || n2 != nil || carry != 0 {
            n3.Next = &ListNode{}
            n3 = n3.Next
        } else {
            break
        }
    }
    return l3
}
