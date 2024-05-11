/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    count := 1
    r1, r2 := head, head
    temp := head
    for r1 != nil {
        if count == k {
            temp = r1
        }
        if r1.Next == nil {
            r2.Val, temp.Val = temp.Val, r2.Val
        }
        if count >= k {
            r2 = r2.Next
        }
        r1 = r1.Next
        count++
    }
    return head
}
