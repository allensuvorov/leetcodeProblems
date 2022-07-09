/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var list3 *ListNode = new(ListNode)
    var head3 *ListNode = list3
    var list1Done, list2Done bool
    
    for !list1Done || !list2Done {
        if !list1Done && list1.Val <= list2.Val {
            list3.Val = list1.Val
            fmt.Println("list1", list1.Val)
            if list1.Next != nil {
                list1 = list1.Next
            } else {
                list1Done = true
                fmt.Println(list1Done)
            }
        } else if !list2Done {
            list3.Val = list2.Val
            fmt.Println("list2", list2.Val)
            if list2.Next != nil{
                list2 = list2.Next
            } else {
                list2Done = true
                fmt.Println(list2Done)
            }
        }
        
        fmt.Println(list1.Val, list2.Val, list3.Val)
        
        if !list1Done || !list2Done {
            list3.Next = new(ListNode)
            list3 = list3.Next    
        }
        
        
    }
    return head3
}
