/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	for ptr := node; ; ptr = ptr.Next {
		ptr.Val = ptr.Next.Val
		if ptr.Next.Next == nil {
			ptr.Next = nil
			break
		}
	}
	return
}