package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [[-2,-1,-1,-1],[]]
	fmt.Println("hello")
	node1 := &ListNode{Val: -2, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1}}}}
	var node2 *ListNode
	h := mergeKLists([]*ListNode{node1, node2})
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type InitHeap []ListNode

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(ListNode))
}

func (h *InitHeap) Pop() any {
	a := *h
	l := len(a)
	item := a[l-1]
	a = a[0 : l-1]
	*h = a
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {

	h := &InitHeap{}
	heap.Init(h)

	for _, node := range lists {
		node := node
		for node != nil {
			heap.Push(h, *node)
			node = node.Next
		}
	}

	newHead := &ListNode{}
	newNode := newHead

	for h.Len() != 0 {

		element := heap.Pop(h)
		node := element.(ListNode)
		newNode.Next = &node
		newNode = newNode.Next
	}
	newNode.Next = nil
	return newHead.Next

	/*
		// algo2
		 if  len(lists) < 1 {
		        return nil
		    }


		    for len(lists) > 1 {
		         var newList []*ListNode
		        for i := 0; i < len(lists); i = i + 2 {
		            l1 := lists[i]

		            var l2 *ListNode
		            if i+1  < len(lists) {
		                l2 = lists[i+1]
		            }


		            mergedLists :=  MergeTwoSortedList(l1, l2)
		            fmt.Println("merged printing is done")
		            newList = append(newList, mergedLists)
		        }
		        lists = newList
		    }

		    return lists[0]
	*/
}

func MergeTwoSortedList(l1, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	newListHead := newList

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			l1 = l1.Next
		} else {
			newList.Next = l2
			l2 = l2.Next
		}
		newList = newList.Next
	}

	if l1 != nil {
		newList.Next = l1
		newList = newList.Next
	}
	if l2 != nil {
		newList.Next = l2
		newList = newList.Next
	}
	return newListHead.Next
}
