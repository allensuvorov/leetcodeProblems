/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

	newHead := head

	for newHead != nil && newHead.Val == val {
		newHead = newHead.Next
	}

	for ptr := newHead; ptr != nil && ptr.Next != nil; ptr = ptr.Next {
		next := ptr.Next
		fmt.Println(ptr, next)
		for next != nil && next.Val == val {
			next = next.Next
		}
		fmt.Println(next)
		if next == nil {
			ptr.Next = nil
			break
		}
		ptr.Next = next
	}
	return newHead
}