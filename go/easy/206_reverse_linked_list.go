/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode; // that dummy takes care of pointing the NEW end node to nil
    for head != nil { // checks that any node exists
        head.Next, prev, head = prev, head, head.Next
    }
    return prev // that dummy is also our new head
}

// full version
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode // dummy
    curr := head
    for curr != nil {
        next := curr.Next // temp
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
