package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    var rev, next *ListNode
    fast, slow := head, head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
		next = slow.Next
        slow.Next = rev
        rev = slow
        slow = next
    }
	if fast != nil { // odd len
		slow = slow.Next
	}
    for rev != nil && rev.Val == slow.Val {
        rev = rev.Next
        slow = slow.Next
    }
    return rev == nil
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
	var testSlice2 []int = []int{1}
	var testSlice3 []int = []int{1,1,2,1}
	var testSlice4 []int = []int{1,2,3}
	fmt.Println(isPalindrome(makeList(testSlice1)))
	fmt.Println(isPalindrome(makeList(testSlice2)))
	fmt.Println(isPalindrome(makeList(testSlice3)))
	fmt.Println(isPalindrome(makeList(testSlice4)))
}