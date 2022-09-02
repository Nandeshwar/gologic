package main

import "fmt"

func main() {
	printArryInReverseOrder([]int{10, 20, 30, 40})
}

func printArryInReverseOrder(a []int) {

	if len(a) == 0 {
		return
	}
	printArryInReverseOrder(a[1:])
	fmt.Println(a[0])
}
