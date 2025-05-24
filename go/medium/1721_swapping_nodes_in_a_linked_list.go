/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    // get list lenth
    // find target nodes
    nodeCount := 0
    var l, r *ListNode // targets
    for runner := head; runner != nil; runner = runner.Next {
        nodeCount++
        // left target node
        if nodeCount == k {
            l = runner
        }

        // right target node
        if nodeCount == k {
            r = head
        }

        // move right target
        if nodeCount > k {
            r = r.Next
        }
    }

    // swap
    l.Val, r.Val = r.Val, l.Val
    return head
}
// nodeCount = 3
// 1 - 2 - 3 - 4 - 5   k = 2
//                     ^
//             r
