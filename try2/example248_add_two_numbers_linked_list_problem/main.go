package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l2 == nil && l1 != nil {
		return l1
	}

	var prev *ListNode
	newList := new(ListNode)
	curr := newList
	c := 0
	for l1 != nil && l2 != nil {
		result := l1.Val + l2.Val + c
		c = result / 10
		curr.Val = result % 10
		prev = curr
		curr.Next = new(ListNode)
		curr = curr.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		result := l1.Val + c
		c = result / 10

		prev = curr
		curr.Val = result % 10
		curr.Next = new(ListNode)
		curr = curr.Next

		l1 = l1.Next
	}

	for l2 != nil {
		result := l2.Val + c
		c = result / 10

		prev = curr
		curr.Val = result % 10
		curr.Next = new(ListNode)
		curr = curr.Next

		l2 = l2.Next
	}

	if c == 0 {
		prev.Next = nil
	} else {
		curr.Val = c
	}

	return newList

}
