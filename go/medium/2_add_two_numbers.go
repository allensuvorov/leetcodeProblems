/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    l3 := &ListNode{} // dummy node
    n1, n2, n3 := l1, l2, l3
    for n1 != nil || n2 != nil || carry == 1{
        sum := carry
        if n1 != nil {
            sum += n1.Val
            n1 = n1.Next
        }
        if n2 != nil {
            sum += n2.Val
            n2 = n2.Next
        }
        carry = sum / 10
        n3.Next = &ListNode{Val:sum % 10}
        n3 = n3.Next
    }
    return l3.Next
}
