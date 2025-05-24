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
    var fromStart, fromEnd *ListNode // targets
    for runner := head; runner != nil; runner = runner.Next {
        nodeCount++
        // left target node
        if nodeCount == k {
            fromStart = runner
        }

        // right target node
        if nodeCount == k {
            fromEnd = head
        }

        // move right target
        if nodeCount > k {
            fromEnd = fromEnd.Next
        }
    }

    // swap
    fromStart.Val, fromEnd.Val = fromEnd.Val, fromStart.Val
    return head
}
// nodeCount = 3
// 1 -> 2 -> 3 -> 4 -> 5   k = 2
//      ^
//      l
//      r
