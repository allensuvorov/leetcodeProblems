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

// raw version, before refactoring

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    // zero nodes
    if head == nil {
        return nil
    }
    // 1 node + // 2 nodes
    if head.Next == nil || head.Next.Next == nil {
        return head
    }
    
    // 3 nodes
    oddHead := head // 1
    oddTail := head
    
    evenHead := head.Next // 2
    evenTail := head.Next
    isOdd := true

    for runner := head.Next.Next; runner != nil; runner = runner.Next {
        if isOdd {
            oddTail.Next = runner // 1-3-5-> nil
            oddTail = oddTail.Next //    t
        } else {
            evenTail.Next = runner   //  2-4->5
            evenTail = evenTail.Next //    t
        }
        isOdd = !isOdd
    }

    evenTail.Next = nil

    oddTail.Next = evenHead
    return oddHead
}
