package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

/*
		10
	  7   9
	5  6

	output:
	bash-3.2$ go run main.go
t.item= 10
t.item= 7
t.item= 5
t.item= 6
t.item= 9
sum= 27
bash-3.2$
*/
func main() {
	five := Tree{item: 5}
	six := Tree{item: 6}
	t := &Tree{item: 10, left: &Tree{item: 7, left: &five, right: &six}, right: &Tree{item: 9}}
	// sum including left and right item
	sum := sumOfItemsBetween(t, 5, 9)
	fmt.Println("sum=", sum)
}

func sumOfItemsBetween(t *Tree, l, r int) int {

	if t == nil {
		return 0
	}
	sum := 0
	fmt.Println("t.item=", t.item)
	if t.item >= l && t.item <= r {
		sum += t.item
	}

	// logic above calculate 1 and
	// below faith

	sum1 := sumOfItemsBetween(t.left, l, r)
	sum2 := sumOfItemsBetween(t.right, l, r)

	return sum + sum1 + sum2
}
