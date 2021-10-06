// https://www.youtube.com/watch?v=OsnikyPMU3Q&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=37
package main

import (
	"container/list"
	"fmt"
)

type Node27 struct {
	item int
	left *Node27
	right *Node27
}
func main() {
	/*
	input: 
	      1
      2 	3 
	33	   4   5
	     7

	output: 7
	*/

	node33 := Node27 {
		item: 33,
	}

	node2 := Node27 {
		item: 2,
		left: &node33,
		
	}
	

	node7 := Node27 {
		item: 7,
	}
	node4 := Node27 {
		item: 4, 
		left: &node7,
	}

	node5 := Node27 {
		item: 5, 
	}

	node3 := Node27 {
		item: 3,
		left: &node4, 
		right: &node5,
	}

	root := Node27 {
		item: 1,
		left: &node2,
		right: &node3,
	}

	

	list := []int{root.item}
	findLeftMostVal(&root, 0, 'r', &list)
	fmt.Println(list)
	fmt.Println(list[len(list)-1])

	root2 := Node27 {
		item: 1,
		left: &node2,
		right: &node3,
	}
	
	fmt.Println(findLeftMostValUsingQueue(&root2))

}

func findLeftMostVal(root *Node27, level int, t rune, list *[]int) {
	if root == nil {
		return
	}

	if t == 'l' {
		if len((*list)) - 1 >= level && (*list)[level] == 0 {
			(*list)[level] = root.item
		} else {
			(*list) = append((*list), root.item)
		}
	}

	findLeftMostVal(root.left, level +1, 'l', list)
	findLeftMostVal(root.right, level +1, 'r', list)
}

func findLeftMostValUsingQueue(root *Node27) int {

	list := list.New()
	list.PushFront(root)

	var node *Node27
	for list.Len() != 0 {
		element := list.Remove(list.Front())
		node = element.(*Node27)
		if node.right != nil {

			list.PushBack(node.right)
		}
		if node.left != nil {

			list.PushBack(node.left)
		}
	}
	return node.item
}


/* output
[1 2 33 4 7]
7
7
*/