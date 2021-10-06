package main

import "container/list"
import "fmt"

func main() {
	s1 := list.New()
	s1.PushBack(1)
	s1.PushBack(2)
	s1.PushBack("3")

	a := s1.Remove(s1.Back())
	fmt.Println(a == 3)

	for e := s1.Front(); e!= nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}