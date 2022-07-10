/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    }
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }    
    var list3 *ListNode = new(ListNode)
    var head3 *ListNode = list3
    
    for {
        if list1.Val <= list2.Val {
            list3.Val = list1.Val
            if list1.Next != nil {
                list1 = list1.Next
            } else {
                list3.Next = list2
                break
            }
        } else {
            list3.Val = list2.Val
            if list2.Next != nil{
                list2 = list2.Next
            } else {
                list3.Next = list1
                break
            }
        }

        list3.Next = new(ListNode)
        list3 = list3.Next
    }
    return head3
}
