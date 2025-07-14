/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getDecimalValue(head *ListNode) int {
    res := 0
    for runner := head; runner != nil; runner = runner.Next {
        res = res * 2
        res = res + runner.Val
    }
    return res
}
