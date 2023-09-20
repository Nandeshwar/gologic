package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

type Pair struct {
	hd   int
	node *Tree
}

func main() {
	/*
			   			10
			   	  5           20
			   2           19   21
			  1  4

			output: 1 2 5 4 10 19 20 21
			Algo:
			   row level traversal - store - horizontal distance and node
			   root horizontal distance = 0
			   left side -1
			   right side +1

			store the result in map :
		      key: hd
		      value: array
	*/

	twentyOne := &Tree{item: 21}
	nineteen := &Tree{item: 19}
	four := &Tree{item: 4}
	two := &Tree{item: 2, right: four, left: &Tree{item: 1}}

	t := new(Tree)
	t.item = 10
	t.left = &Tree{item: 5, left: two}
	t.right = &Tree{item: 20, left: nineteen, right: twentyOne}

	arr := verticalTraversal(t)
	fmt.Println("arr=", arr)
}

func verticalTraversal(t *Tree) []int {
	var result []int
	m := make(map[int][]int)
	pair := &Pair{hd: 0, node: t}
	q := list.New()
	q.PushBack(pair)

	for q.Len() != 0 {
		fmt.Println("q.Len()=", q.Len())
		rowLen := q.Len()
		for rowLen != 0 {
			element := q.Remove(q.Front())
			pair := element.(*Pair)
			hd := pair.hd
			node := pair.node
			rowLen--

			if a, ok := m[hd]; ok {
				a = append(a, node.item)
				m[hd] = a
			} else {
				m[hd] = []int{node.item}
			}

			if node.left != nil {
				newPair := &Pair{
					hd:   hd - 1,
					node: node.left,
				}

				q.PushBack(newPair)
			}

			if node.right != nil {
				newPair := &Pair{
					hd:   hd + 1,
					node: node.right,
				}

				q.PushBack(newPair)
			}

		}
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, m[k]...)
	}
	return result
}
