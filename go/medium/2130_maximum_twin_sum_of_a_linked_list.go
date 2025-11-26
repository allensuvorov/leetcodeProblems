/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// reverse 1st half
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


// reverse 1st half
func pairSum(head *ListNode) int {
    // space O(1)

    // read len and reverse first half
    fast := head // for finding the middle
    head2 := head // reversed new head2
    
    var head1 *ListNode // dummy, to be the head1
    
    for fast != nil && fast.Next != nil {
        // end of list runner
        fast = fast.Next.Next

        // reversal mechanics
        temp := head2.Next
        head2.Next = head1 // reversal operation
        head1 = head2
        head2 = temp
    }

    // fmt.Println(head1.Val, head2.Val)
    // get result
    maxSum := 0
    for head1 != nil && head2 != nil {
        fmt.Println(head1.Val, head2.Val)
        maxSum = max(maxSum, head1.Val + head2.Val)
        head1 = head1.Next
        head2 = head2.Next
    }
    return maxSum
}
