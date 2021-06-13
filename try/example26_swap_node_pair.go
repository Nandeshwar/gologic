// https://www.youtube.com/watch?v=-xwX521Ija4&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=36
package main

import (
	"fmt"
)

type Node26 struct {
	item int
	next *Node26
}

func main() {
	root := &Node26{item: 1, next: &Node26{item:2, next: &Node26{item: 3, next: &Node26{item:4, next: &Node26{item: 5, next: &Node26{item:6}}}}}}
	
	
	printNode26(root)
	newRoot := exchangeNode(root)
	fmt.Println("After exchanging nodes")
	printNode26(newRoot)

	root2 := &Node26{item: 1, next: &Node26{item:2, next: &Node26{item: 3, next: &Node26{item:4, next: &Node26{item: 5, next: &Node26{item:6}}}}}}
	fmt.Println("After using 2nd logic")
	newRoot2 := exchangeNode2(root2)
	printNode26(newRoot2)

}

func printNode26(root *Node26) {
	if root == nil {
		return
	}
	fmt.Println(root.item)
	printNode26(root.next)
}

func exchangeNode(root *Node26) *Node26 {

	current := root
	p := current
	if root.next != nil {
		root = root.next
	}

	// current node skip 1 node in middle and p will point node which is skipped i.e middle for node connection
	// current node is 1st node then 3rd node(1, 3) for 1, 2, 3, 4
	// p will current node 1 here and it's next will point to current->next ie. 1 -> 4 after node exchage 2, 1, 4, 3
	for current != nil {
		next := current.next
		tmp := next.next
		next.next = current
		current.next = tmp
		p = current
		current = current.next
		if current != nil {
			p.next = current.next
		}
	}

	return root
}

// create tmp node and point it to 1st node
// currentNode = tmp
// firstNode = currentNode-> next
// seoncNode = currentNode->next->next
// exchange node 
// set currentNode = currentNode->next->next i.e 2nd node so 3 and 4 node can be treated as 1st and and 2nd node do the exchanges.

func exchangeNode2(root *Node26) *Node26 {
	tmp := &Node26{
		item:0,
		next: root,
	}

	current := tmp

	for current.next != nil && current.next.next != nil {
		first := current.next
		second := current.next.next

		first.next = second.next
		second.next = first
		current.next = second
		current = current.next.next
	}
	return tmp.next
}
