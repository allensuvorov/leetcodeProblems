/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    listLenth := 0
    for runner := head; runner != nil; runner = runner.Next {
        listLenth++
    }

    second := splitList(head, listLenth / 2)
    second = reverseList(second)

    // find max
    first := head
    maxPairSum := 0
    for range listLenth / 2 {
        maxPairSum = max(maxPairSum, first.Val + second.Val)
        first = first.Next
        second = second.Next
    }
    return maxPairSum
}

func splitList(head *ListNode, half int) *ListNode {
    for range half {
        head = head.Next
    }
    return head
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for cur := head; cur != nil; cur.Next, prev, cur = prev, cur, cur.Next {}
    return prev
}
