package main

import "fmt"

type Node struct {
	item   int
	next   *Node
	random *Node
}

func (n *Node) display() {
	head := n
	for head != nil {
		item := head.item
		var nextItem int
		var randomItem int

		if head.next != nil {
			nextItem = head.next.item
		}
		if head.random != nil {
			randomItem = head.random.item
		}
		fmt.Println("val=", item, "next.val=", nextItem, "random.val=", randomItem)
		head = head.next
	}
}

func main() {
	five := &Node{item: 5}
	four := &Node{item: 4, next: five}
	three := &Node{item: 3, next: four}
	two := &Node{item: 2, next: three}
	one := &Node{item: 1, next: two, random: three}
	two.random = one
	four.random = two
	five.random = four
	three.random = three

	one.display()

	m := make(map[*Node]*Node)
	newNode := createDuplicate(one, m)
	fmt.Println("New Node...")
	newNode.display()

	fmt.Println("new logic. order of n and no space")
	duplicateNodes := createDuplicate2(newNode, m)
	fmt.Println("Duplicate nodes created by algorithm two O(n) time complexity and no space ")
	duplicateNodes.display()
}

// time complexity: O(n), space complexity O(n)
func createDuplicate(head *Node, m map[*Node]*Node) *Node {
	curr := head
	for curr != nil {
		newNode := &Node{item: curr.item}
		m[curr] = newNode
		curr = curr.next
	}

	curr = head

	for curr != nil {
		newNode := m[curr]
		newNode.next = m[curr.next]
		newNode.random = m[curr.random]
		curr = curr.next
	}

	return m[head]

}

// time complexity O(n) - no space used
func createDuplicate2(head *Node, m map[*Node]*Node) *Node {
	// create duplicate nodes
	/*
	 1   2   3    4   5   nil
	  \ / \ / \  / \ / \ /
	   1   2   3   4   5
	*/
	var dupHead *Node
	i := 0
	curr := head
	for curr != nil {
		currNext := curr.next

		newNode := &Node{item: curr.item}
		curr.next = newNode
		newNode.next = currNext

		curr = currNext
		if i == 0 {
			dupHead = newNode
			i++
		}
	}
	fmt.Println("Display currently created structure with just duplicate node.")
	head.display()

	fmt.Println("Now let's assign random pointer and display")
	curr = head
	for curr != nil {
		curr.next.random = curr.random.next // understand it by drawing image of nodes
		curr = curr.next.next
	}
	head.display()

	// Now separate dulicate records with original records

	curr = head
	for curr != nil {
		nextCurr := curr.next.next

		originalNext := curr.next.next
		var dupNext *Node
		if originalNext != nil {
			dupNext = originalNext.next
		}

		curr.next.next = dupNext
		curr.next = originalNext

		curr = nextCurr
	}

	return dupHead
}

/*
val= 1 next.val= 2 random.val= 3
val= 2 next.val= 3 random.val= 1
val= 3 next.val= 4 random.val= 3
val= 4 next.val= 5 random.val= 2
val= 5 next.val= 0 random.val= 4
New Node...
val= 1 next.val= 2 random.val= 3
val= 2 next.val= 3 random.val= 1
val= 3 next.val= 4 random.val= 3
val= 4 next.val= 5 random.val= 2
val= 5 next.val= 0 random.val= 4
new logic. order of n and no space
Display currently created structure with just duplicate node.
val= 1 next.val= 1 random.val= 3
val= 1 next.val= 2 random.val= 0
val= 2 next.val= 2 random.val= 1
val= 2 next.val= 3 random.val= 0
val= 3 next.val= 3 random.val= 3
val= 3 next.val= 4 random.val= 0
val= 4 next.val= 4 random.val= 2
val= 4 next.val= 5 random.val= 0
val= 5 next.val= 5 random.val= 4
val= 5 next.val= 0 random.val= 0
Now let's assign random pointer and display
val= 1 next.val= 1 random.val= 3
val= 1 next.val= 2 random.val= 3
val= 2 next.val= 2 random.val= 1
val= 2 next.val= 3 random.val= 1
val= 3 next.val= 3 random.val= 3
val= 3 next.val= 4 random.val= 3
val= 4 next.val= 4 random.val= 2
val= 4 next.val= 5 random.val= 2
val= 5 next.val= 5 random.val= 4
val= 5 next.val= 0 random.val= 4
Duplicate nodes created by algorithm two O(n) time complexity and no space
val= 1 next.val= 2 random.val= 3
val= 2 next.val= 3 random.val= 1
val= 3 next.val= 4 random.val= 3
val= 4 next.val= 5 random.val= 2
val= 5 next.val= 0 random.val= 4

*/
