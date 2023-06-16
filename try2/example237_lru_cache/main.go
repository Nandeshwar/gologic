package main

/*
1. Remove least access item
2. access item will be at first

Algorithm
----------
put(1, 1) : 1 will in front
put(2,2): 2 will in front
get(1,1) : 1 will in front


*/

import (
	"container/list"
	"fmt"
)

type Pair struct {
	element *list.Element
	value   int
}

type Cache struct {
	m    map[int]Pair
	q    *list.List
	size int
}

func main() {
	c := New(5)
	c.put(1, 1)
	fmt.Println(c.get(1))
	c.put(2, 2)
	fmt.Println(c.get(2))

	fmt.Println(c.get(1))
	fmt.Println("front=", c.q.Front().Value)
}

func New(size int) *Cache {
	return &Cache{m: make(map[int]Pair), q: list.New(), size: size}
}

func (c Cache) put(key, value int) {
	if c.q.Len() == c.size {
		c.q.Remove(c.q.Back())

	}
	pair, ok := c.m[key]
	if ok {
		frontElement := c.q.Front()
		c.q.MoveBefore(pair.element, frontElement)

	} else {
		e := c.q.PushFront(key)
		c.m[key] = Pair{element: e, value: value}
	}

}

func (c Cache) get(key int) int {

	pair, ok := c.m[key]
	if !ok {
		return -1
	}

	c.q.MoveBefore(pair.element, c.q.Front())

	return pair.value
}
