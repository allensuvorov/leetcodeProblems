/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil {
        return head
    }
    var cur, first, last *ListNode
    var len int
    // get the length
    for cur = head ; cur != nil; cur = cur.Next {
        len++
        last = cur
    }
    
    if k % len == 0 {
        return head
    }

    if k > len {
        k = k % len
    }

    cur = head
    for i := 1; i < len - k; i ++ {
        cur = cur.Next
    }
    first = cur.Next
    cur.Next = nil
    last.Next = head

    return first
}
