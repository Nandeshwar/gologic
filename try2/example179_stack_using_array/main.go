package main

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
