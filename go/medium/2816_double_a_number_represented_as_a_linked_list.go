/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    if head.Val > 4 {
        dummy := &ListNode{Next: head}
        head = dummy
    }

    for runner := head; runner != nil; runner = runner.Next {
        carry := 0
        if runner.Next != nil && runner.Next.Val > 4 {
            carry = 1
        }
        runner.Val = runner.Val * 2 + carry
        runner.Val = runner.Val % 10
    }
    return head
}
