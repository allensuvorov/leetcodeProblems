/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    minDist := 100000
    maxDist := -1
    curDist := -1
    for prev, cur := head, head.Next; cur.Next != nil; prev, cur = cur, cur.Next {
        if isCritical(prev.Val, cur.Val, cur.Next.Val) {
            if curDist != -1 {
                minDist = min(minDist, curDist)
            }
            curDist = 1
        } else {
            if curDist != -1 {
                curDist++
            }
        }
    }

    curDist = -1
    for prev, cur := head, head.Next; cur.Next != nil; prev, cur = cur, cur.Next {
        if isCritical(prev.Val, cur.Val, cur.Next.Val) {
            if curDist != -1 {
                maxDist = max(maxDist, curDist)
                curDist++
            } else {
                curDist = 1
            }
        } else {
            if curDist != -1 {
                curDist++
            }
        }
    }

    if maxDist == -1 {
        minDist = -1
    }
    return []int{minDist, maxDist}
}

func isCritical(a, b, c int) bool {
    return (a < b && b > c) || (a > b && b < c)
}
