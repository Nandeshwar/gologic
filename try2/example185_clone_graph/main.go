package main

import (
	"fmt"
)

type Node struct {
	item       int
	neighbours []*Node
}

func (n *Node) display() {
	fmt.Println("------------Begin_________")
	fmt.Println("item=", n.item)
	fmt.Println("Neighbours")
	for _, v := range n.neighbours {
		fmt.Println(v.item)
	}
	fmt.Println("------------End_________")
}

func main() {
	/*
		1 ---------- 2
		|            |
		|            |
		|            |
		4 ---------- 3
	*/
	three := &Node{item: 3}
	four := &Node{item: 4}
	two := &Node{item: 2}
	oneNeighbours := []*Node{two, four}
	one := &Node{item: 1, neighbours: oneNeighbours}
	two.neighbours = []*Node{one, three}
	three.neighbours = []*Node{two, four}
	four.neighbours = []*Node{one, three}

	two.display()
	one.display()
	three.display()
	four.display()

	newNode := cloneGraph(one)
	fmt.Println("**************After cloning*******************")
	newNode.display()
	for _, v := range newNode.neighbours {
		v.display()
	}

}

func cloneGraph(root *Node) *Node {
	m := make(map[*Node]*Node)
	return clone(root, m)
}

func clone(node *Node, m map[*Node]*Node) *Node {

	exitingNode, ok := m[node]
	if ok {
		return exitingNode
	}
	newNode := &Node{item: node.item}
	m[node] = newNode

	for _, n := range node.neighbours {
		newNode.neighbours = append(newNode.neighbours, clone(n, m))
	}
	return newNode

}
