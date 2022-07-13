package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    // walker reverses the first half of the list
    // when runner reaches finish, walker points to center or pre-center node
    // 123 123 // then we get two lists, and can compare values from the begining.
    
    var cur, prev, next, runner, list1, list2 *ListNode
    runner = head
    cur = head
    
	if head.Next == nil {
        return true
    }
    if head.Next.Next == nil {
        return head.Val == head.Next.Val
    }
    if head.Next.Next.Next == nil {
        return head.Val == head.Next.Next.Val
    }
    // reversing 1 - n/2
    for head != nil {
        runner = runner.Next.Next
        next = cur.Next 
        cur.Next = prev
        prev = cur
        cur = next
        
        // fmt.Println(runner)
        if runner.Next == nil { // odd len
            list2 = next.Next
            // fmt.Println(list2)
            break
        }
        if runner.Next.Next == nil { // even len
            list2 = next
            // fmt.Println(list2)
            break
        }
    }
    
    list1 = prev
    
    // comparing list1 and list2
	fmt.Println(list1.Val,list2.Val)
    for list1 != nil && list2 != nil {
        if list1.Val != list2.Val {
            return false
        }
        list1 = list1.Next
        list2 = list2.Next
    }
    return true
    
}

func makeList(nums []int) *ListNode {
	fmt.Println(nums)
	var head *ListNode = new(ListNode)
	var cur *ListNode = head
	for i:=0; i<len(nums)-1; i++{
		cur.Val = nums[i]
		var next *ListNode = new(ListNode)
		cur.Next = next
		cur = next 
	}
	cur.Val = nums[len(nums)-1]
	cur.Next = nil
	return head
}

func main() {
	var testSlice1 []int = []int{1,2,3,4,3,2,1}
	fmt.Println(isPalindrome(makeList(testSlice1)))
}
