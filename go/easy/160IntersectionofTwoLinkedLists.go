/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a1, a2 := headA, headA
	b1, b2 := headB, headB

	for a1 != nil && b1 != nil {
		if b1.Next != nil && a1.Next == nil {
			b2 = b2.Next
		}
		if a1.Next != nil && b1.Next == nil {
			a2 = a2.Next
		}
		if a1.Next != nil {
			a1 = a1.Next
		}
		if b1.Next != nil {
			b1 = b1.Next
		}
		if a1.Next == nil && b1.Next == nil {
			break
		}
	}
	for a1 == b1 {
		fmt.Println(a2.Val, b2.Val)
		if a2 == b2 {
			fmt.Println("intersection", a2.Val, b2.Val)
			return a2
		}
		if a2.Next == nil {
			fmt.Println("break", a2.Val, b2.Val)
			break
		}
		a2 = a2.Next
		b2 = b2.Next
	}
	return nil
}

// the nice and consise solution
func getIntersectionNodeA(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}