package main

import "fmt"

type Node struct {
	item int
	next *Node
}

func main() {
	one := &Node{1, nil}
	otherTwo := &Node{2, one}
	two := &Node{2, otherTwo}
	head := &Node{0, two}

	slow := head
	fast := head
 
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	palindrome := true
	// slow will reach to 2nd half and reverse that
	secondHalf := reverse(slow)
	
	// now compare 1st item of head with 1st item of 2nd half
	for secondHalf != nil {
		if secondHalf.item != head.item {
			palindrome = false
			break
		}
		secondHalf = secondHalf.next
		head = head.next
	}

	if palindrome {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}

}

func reverse(head *Node) *Node {
	var prev *Node
	for head != nil {
		tmp := head.next
		head.next = prev
		prev = head
		head = tmp
	}
	return prev
}
