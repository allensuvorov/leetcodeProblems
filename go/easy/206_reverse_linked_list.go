/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var temp *ListNode
    for head != nil {
        temp, head.Next = head.Next, temp
        temp, head = head, temp
    }
    return temp
}
