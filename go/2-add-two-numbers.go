/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    l3 := &ListNode{}
    n1, n2, n3 := l1, l2, l3

    for {
        var v1, v2 int
        if n1 != nil {
            v1 = n1.Val
            n1 = n1.Next
        } 

        if n2 != nil {
            v2 = n2.Val
            n2 = n2.Next
        }

        sum := v1 + v2 + carry
        n3.Val = sum % 10
        if sum > 9 {
            carry = 1
        } else {
            carry = 0
        }

        if n1 != nil || n2 != nil || carry != 0 {
            n3.Next = &ListNode{}
            n3 = n3.Next
        } else {
            break
        }
    }
    return l3
}
