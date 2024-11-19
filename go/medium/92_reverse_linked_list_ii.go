/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    i := 1
    before := &ListNode{Next:head}
    for runner := head; runner != nil && i < left; runner = runner.Next{
        before = runner
        i++
    }
    // before, segmentHead
    segmentHead := before.Next

    var prev *ListNode
    cur := segmentHead

    for cur != nil && i <= right {
        cur. Next, prev, cur = prev, cur, cur.Next
        i++
    }
    // reconnect reversed segment
    before.Next = prev // prev is new segmentHead
    segmentHead.Next = cur // cur is after
    
    // edge case
    if left == 1 {
        head = prev
    }
    return head
}
