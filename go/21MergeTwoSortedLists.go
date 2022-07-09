/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    var list3 *ListNode = new(ListNode)
    
    var node1 ListNode = *list1
    var node2 ListNode = *list2
    var node3 ListNode = *list3
    
    for node1.Next != nil && node2.Next != nil {
        
        if node1.Val >= node2.Val {
            node3.Val = node1.Val
            node1 = *node1.Next
        } else {
            node3.Val = node2.Val
            node2 = *node2.Next
        }
        
        fmt.Println( node1, node2, node3)
        
        node3.Next = new(ListNode)
        
        fmt.Println("head is", list3)
        
        node3 = *node3.Next
        fmt.Println("head is node3", list3)
    }
    
    if node1.Val >= node2.Val {
            node3.Val = node1.Val
    } else {
            node3.Val = node2.Val
    }
    
    
    // fmt.Println(list1, node1, node2, node3, list3)
    
    return list3
}
