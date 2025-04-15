/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    // find GCD - iteration from min(a,b) .... 1
    getGCD := func(a, b int) int {
        for x := min(a, b); x > 0; x-- {
            if a % x == 0 && b % x == 0 {
                return x
            }
        }
        return 1
    }

    // insert in SLL
    for a, b := head, head.Next; b != nil; b = b.Next {
        gcd := &ListNode{getGCD(a.Val, b.Val), b}
        a.Next = gcd
        a = b
    }
    return head
}
