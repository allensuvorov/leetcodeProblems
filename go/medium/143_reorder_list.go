/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type StackNode struct{
    Val *ListNode
    Next *StackNode
}

func reorderList(head *ListNode)  {
    cur := head
    var stack *StackNode
    // fill up stack
    for cur != nil {
        newStackNode := StackNode{cur, stack}
        stack = &newStackNode
        cur = cur.Next
    }
    fmt.Println("Stack is ", stack.Val.Val)
    cur = head
    for cur != stack.Val && cur.Next != stack.Val {
        next := cur.Next
        cur.Next = stack.Val // 1 -> 4
        stack = stack.Next
        cur = cur.Next
        cur.Next = next // 4 -> 2
        cur = cur.Next
        if cur == stack.Val {
            cur.Next = nil    
            break
        }
        if cur.Next == stack.Val {
            cur.Next.Next = nil
            break
        }

    }
}
