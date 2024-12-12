/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    switch n {
    case 0:
        return nil
    case 1:
        return lists[0]
    }

    return mergeTwoLists(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    runner := dummy
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            runner.Next = list1
            list1 = list1.Next
        } else {
            runner.Next = list2
            list2 = list2.Next
        }
        runner = runner.Next
    }
    
    if list1 == nil {
        runner.Next = list2
    } else {
        runner.Next = list1
    }
    return dummy.Next
}
