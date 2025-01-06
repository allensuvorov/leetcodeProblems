/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    var dummy *ListNode
    prev := dummy
    cur := head

    for fast := head; fast != nil && fast.Next != nil; {
        cur.Next, prev, cur, fast = prev, cur, cur.Next, fast.Next.Next
    }

    res := 0
    for r1, r2 := prev, cur; r1 != nil; r1, r2 = r1.Next, r2.Next {
        res = max(res, r1.Val + r2.Val)
    }
    return res
}
