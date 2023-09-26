/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    // edge cases
    if head.Next == nil || left == right{
        return head
    }

    // init end of first slice
    r1 := head
    for i := 1; i + 1 < left; i++ {
        r1 = r1.Next
    }
    
    // reverse slice
    cur := r1
    if left != 1 {
        cur = r1.Next
    }

    r2 := cur // future end of second (reversed) slice
    var temp *ListNode
    for i := 1; i <= right - left + 1 && cur != nil; i++ {
        temp, cur.Next = cur.Next, temp
        cur, temp = temp, cur
    }
    l2 := temp

    // insert reversed slice
    if left != 1 {
        r1.Next = l2
    } else {
        head = l2
    }

    if cur != nil {
        r2.Next = cur
    }
    return head
}
