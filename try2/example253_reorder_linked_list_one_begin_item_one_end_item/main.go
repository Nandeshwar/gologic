package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Input: head = [1,2,3,4]
Output: [1,4,2,3]

// slow and fast pointer technique
// reverse 2nd portion
//
*/
func main() {

}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	slow := head
	fast := head

	var prev *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == nil || fast.Next == nil && prev != nil {
			prev.Next = nil
		}

		prev = slow
	}

	head2 := reverse(slow)

	//var prevHead *ListNode
	for head != nil && head2 != nil {
		tmp1 := head.Next
		tmp2 := head2.Next

		head.Next = head2
		head = head.Next

		head.Next = tmp1

		if head.Next != nil {
			head = head.Next
		}

		head2 = tmp2

	}

	if head2 != nil {
		head.Next = head2
	}

}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

func print(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)

	print(head.Next)
}
