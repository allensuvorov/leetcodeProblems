/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    
    var cur, prev, next *ListNode
    cur = head
    for cur != nil{
        next = cur.Next // address of (2)
        cur.Next = prev // (1)->(nil) pointer
        prev = cur // save address of (1)
        cur = next // go to (2)
    }
    return prev
}
