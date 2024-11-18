/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    nodes := []*ListNode{}
    for runner := head; runner != nil; runner = runner.Next {
        nodes = append(nodes, runner)
    }

    slices.SortFunc(nodes, func(a, b *ListNode) int {
        return a.Val - b.Val
    })

    nodes[len(nodes)-1].Next = nil
    for i := 0; i < len(nodes) - 1; i++ {
        nodes[i].Next = nodes[i+1]
    }

    return nodes[0]
}
