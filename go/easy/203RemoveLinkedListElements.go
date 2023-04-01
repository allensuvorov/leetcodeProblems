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

		for next != nil && next.Val == val {
			next = next.Next
		}

		if next == nil {
			ptr.Next = nil
			break
		}
		if ptr.Next != next {
			ptr.Next = next
		}
	}
	return newHead
}