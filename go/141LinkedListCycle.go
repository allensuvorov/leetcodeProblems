/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    
    for walker, runner := head, head.Next; 
        runner.Next != nil && runner.Next.Next != nil; 
        walker, runner = walker.Next, runner.Next.Next {
          
        if walker == runner {
            return true
        }  
    }
    
    return false
}
