package main

func main() {

}

// slow and fast pointer technique
// move fast to n
// then move slow and fast by 1
// corner case also count number of nodes total if slow did not move and want to remove 1st item:
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	slow := head
	fast := head

	nodeCnt := 1

	for i := 0; i < n && fast != nil && fast.Next != nil; i++ {
		fast = fast.Next
		nodeCnt++
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		nodeCnt++
	}
	return head
}
