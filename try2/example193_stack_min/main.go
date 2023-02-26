package main

import (
	"container/list"
	"fmt"
)

type Stack1 struct {
	s1 *list.List
	s2 *list.List
}

func New1() *Stack1 {
	return &Stack1{s1: list.New(), s2: list.New()}
}

func main() {
	stack1 := New1()
	stack1.Push(10)
	fmt.Println(stack1.Min())
	stack1.Push(20)
	fmt.Println(stack1.Min())
	stack1.Push(1)
	fmt.Println(stack1.Min())

	stack1.Pop()
	fmt.Println(stack1.Min())

	fmt.Println("*********Algorithm2....")
	stack2 := New2()
	stack2.Push(10)
	fmt.Println(stack2.Min())
	stack2.Push(20)
	fmt.Println(stack2.Min())
	stack2.Push(1)
	fmt.Println(stack2.Min())

	stack1.Pop()
	fmt.Println(stack1.Min())

}

func (s *Stack1) Push(item int) {
	if s.s1.Len() == 0 {
		s.s1.PushBack(item)
		s.s2.PushBack(item)
	} else {
		element := s.s2.Back().Value
		sItem := element.(int)

		s.s1.PushBack(item)
		s.s2.PushBack(min(item, sItem))
	}
}

func (s *Stack1) Pop() int {
	if s.s1.Len() == 0 {
		return -1
	}
	element := s.s1.Remove(s.s1.Back())
	item := element.(int)

	s.s2.Remove(s.s2.Back())
	return item
}

func (s *Stack1) Min() int {
	if s.s2.Len() == 0 {
		return -1
	}

	element := s.s2.Back().Value
	item := element.(int)
	return item
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Pair struct {
	item    int
	minItem int
}

type Stack2 struct {
	stack *list.List
}

func New2() *Stack2 {
	return &Stack2{list.New()}
}

func (s Stack2) Push(item int) {
	if s.stack.Len() == 0 {
		s.stack.PushBack(Pair{item, item})
		return
	}

	element := s.stack.Remove(s.stack.Back())
	sItem := element.(Pair)

	s.stack.PushBack(Pair{item, min(item, sItem.minItem)})
}

func (s Stack2) Pop() int {
	if s.stack.Len() == 0 {
		return -1
	}

	element := s.stack.Remove(s.stack.Back())
	item := element.(Pair)
	return item.item
}

func (s Stack2) Min() int {
	if s.stack.Len() == 0 {
		return -1
	}

	element := s.stack.Back().Value
	item := element.(Pair)
	return item.minItem
}
