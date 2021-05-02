// https://www.youtube.com/watch?v=Ds5e1A88j7Q&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=7
package main

import "fmt"

type TreeNode struct {
	Item int
	Children []TreeNode
}

type Stack struct {
	items []TreeNode
	ind int
}


func main() {
	stack := &Stack{
		ind: -1,
		items: make([]TreeNode, 10),
	}

	fmt.Println(stack.pop())
	node := TreeNode{
		Item: 1,
		Children: []TreeNode{
			{Item: 3, Children: []TreeNode{
					{Item: 5}, 
					{Item: 6},
				},
			},
			{Item: 2},
			{Item: 4},
		},
	}

	resultStack := Stack{ind: -1, items: make([]TreeNode, 10)}

	stack.push(node)
	for ; stack.ind != -1; {
		node := stack.pop()
		//fmt.Println(node.Item)

		resultStack.push(*node)
		
		for _, node := range node.Children {
			stack.push(node)
		}
	}

	// Display result
	for ; resultStack.ind != -1; {
		node := resultStack.pop()
		fmt.Println(node.Item)
	}


}

func(s *Stack) push(item TreeNode) {
	if s.ind == 10 {
		fmt.Println("Stack Overflow")
		return
	} 

	s.ind++
	s.items[s.ind] = item
}

func(s *Stack) pop() *TreeNode {
	if s.ind == -1 {
		fmt.Println("Stack is empty")
		return nil
	}
	
	item := s.items[s.ind]
	s.ind--
	return &item
}

