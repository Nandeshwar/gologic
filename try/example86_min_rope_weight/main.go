// priority queue example
// youtube: https://www.youtube.com/watch?v=Eb1A6nm_Nic&list=PLUcsbZa0qzu3yNzzAxgvSgRobdUUJvz7p&index=35
package main

import (
	"fmt"

	"github.com/x1m3/priorityQueue"
)

type Item int

func (i Item) HigherPriorityThan(other priorityQueue.Interface) bool {
	return i < other.(Item)
}

func min_rope_weight(a []int) int {
	list := priorityQueue.New()
	for _, v := range a {
		list.Push(Item(v))
	}

	totalSum := 0

	for {
		r := list.Pop()

		val1 := r.(Item)

		r = list.Pop()
		if r == nil {
			break
		}
		val2 := r.(Item)

		fmt.Println("val1=", val1, "val2", val2)

		sum := int(val1) + int(val2)
		totalSum = totalSum + sum
		fmt.Println("sum", sum)
		list.Push(Item(sum))

	}
	return totalSum

}
