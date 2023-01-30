package main

/*
	use of stack:
	  implement queue using 2 stacks:
	    3
		2
		1
		
		  ->   get from stack1 and put in stack2
				     1
		             2
		             3
					
					remove from stack
					
					
	use of queue:
	    implement stack using two queues
		1 
		   q1
		       
			
			if two or more than two times
			  move 1 to q2 and insert next item(2 here) in q1 and move all items from q2 to q1
			  1
				    to q2
					
		     put 2 to q1
			     move items from q2
		   q1: 2, 1
*/

import (
	"fmt"
)

type Stack struct {
	items []int
	size  int
	top   int
}

func New(size int) *Stack {
	return &Stack{items: make([]int, size), size: size, top: -1}
}

func main() {
	s := New(3)
	fmt.Println(s.size)

	s.Push(10)
	fmt.Println(s.items)
	s.Push(20)
	fmt.Println(s.items)
	s.Push(30)
	fmt.Println(s.items)
	s.Push(40)
	fmt.Println(s.items)
	fmt.Println(s.Pop())
	fmt.Println(s.items)

	s.Push(40)
	fmt.Println(s.items)
	fmt.Println(s.Pop())
	fmt.Println(s.items)
	fmt.Println(s.Pop())
	fmt.Println(s.items)
	fmt.Println(s.Pop())
	fmt.Println(s.items)
	fmt.Println(s.Pop())
	fmt.Println(s.items)
	s.Push(40)
	fmt.Println(s.items)
	fmt.Println(s.Pop())
	fmt.Println(s.items)
}

func (s *Stack) Push(item int) bool {
	if s.top == s.size-1 {
		fmt.Println("Stack is full")
		return false
	}

	s.top++
	s.items[s.top] = item
	return true
}

func (s *Stack) Pop() int {
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return -1
	}

	item := s.items[s.top]
	s.top--
	return item
}
