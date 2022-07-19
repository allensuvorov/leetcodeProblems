/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var dec, l float64
	// get len
	ptr := head
	for ptr != nil {
		ptr = ptr.Next
		l++
	}
	ptr = head
	for i := l - 1; i >= 0; i-- {
		dec += float64(ptr.Val) * math.Pow(2, i)
		fmt.Println(ptr.Val, math.Pow(2, i))
		ptr = ptr.Next
	}
	return int(dec)
}

// sweet solution to multiply and add next val
func getDecimalValue(head *ListNode) int {
	var num = head.Val

	for head.Next != nil {
		num = num*2 + head.Next.Val
		fmt.Println(num, head.Next.Val)
		head = head.Next
	}
	return num
}