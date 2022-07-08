/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var list3 ListNode
    
    if list1.Val >= list2.Val {
        list3.Val = list1.Val
    }
    
    // type ListNode struct {
    //     Val int
    //     Next *ListNode
    // }
    
    fmt.Println(list3.Val)
    return list1
}
