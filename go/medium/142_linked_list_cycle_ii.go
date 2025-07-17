/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    cycle := false

    for slow, fast := head, head.Next; fast != nil && fast.Next != nil; {
        
        if slow == fast {
            cycle = true
            break
        }
        slow, fast = slow.Next, fast.Next.Next
    }
    
    if cycle {
        seen := make(map[*ListNode]bool)
        for runner := head; !seen[runner]; runner = runner.Next {
            seen[runner] = true
            if seen[runner.Next] {
                return runner.Next
            }
        }
    }

    return nil
}
   
