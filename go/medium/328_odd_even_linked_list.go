/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    odd, even := &ListNode{}, &ListNode{}
    evenDummy := even
    isEven := false
    for runner := head; runner != nil; runner = runner.Next {
        if isEven {
            even.Next = runner
            even = even.Next
        } else {
            odd.Next = runner
            odd = odd.Next
        }
        isEven = !isEven
    }
    even.Next = nil
    odd.Next = evenDummy.Next

    return head
}
