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
    curMin := -1
    curMax := -1
    for prev, cur := head, head.Next; cur.Next != nil; prev, cur = cur, cur.Next {
        if isCritical(prev.Val, cur.Val, cur.Next.Val) {
            if curMin != -1 {
                minDist = min(minDist, curMin)
            }
            curMin = 1
            if curMax != -1 {
                maxDist = max(maxDist, curMax)
                curMax++
            } else {
                curMax = 1
            }
        } else {
            if curMin != -1 {
                curMin++
            }
            if curMax != -1 {
                curMax++
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
