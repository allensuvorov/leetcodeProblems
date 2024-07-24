/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    if head.Val > 4 {
        head = &ListNode{Next: head}
    }

    for runner := head; runner != nil; runner = runner.Next {
        runner.Val = (runner.Val * 2) % 10
        if runner.Next != nil && runner.Next.Val > 4 {
            runner.Val += 1
        }
    }
    return head
}
