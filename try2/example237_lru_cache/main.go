package main

/*
1. Remove least access item
2. access item will be at first

Algorithm
----------
put(1, 1) : 1 will in front
put(2,2): 2 will in front
get(1,1) : 1 will in front


Java:
  LinkedHashMap can be used
  - item will be inserted in front and insertion order will be preserved.

*/

import (
	"container/list"
	"fmt"

	"github.com/emirpasic/gods/maps/linkedhashmap"
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

type Cache2 struct {
	m *linkedhashmap.Map
}

func main() {
	c := New(5)
	c.put(1, 1)
	fmt.Println(c.get(1))
	c.put(2, 2)
	fmt.Println(c.get(2))

	fmt.Println(c.get(1))
	fmt.Println("front=", c.q.Front().Value)

	c2 := New2()
	c2.put(1, 1)
	c2.put(2, 2)
	c2.put(3, 3)
	fmt.Println(c2.get(1))
	fmt.Println(c2.m)

	fmt.Println("Algo 2......")
	it := c2.m.Iterator()

	for it.Next() {
		fmt.Println(it.Value())
	}

	fmt.Println(c2.m.Iterator())

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

func New2() *Cache2 {
	return &Cache2{m: linkedhashmap.New()}
}

func (c Cache2) put(key, value int) {

	_, ok := c.m.Get(key)
	if ok {
		c.m.Remove(key)
	}
	c.m.Put(key, value)
}

func (c Cache2) get(key int) int {
	value, ok := c.m.Get(key)
	if ok {
		c.m.Remove(key)
		c.m.Put(key, value)
	}

	return value.(int)
}
