package main

import "fmt"

func main() {
	printAsc(5)
	fmt.Println("--------")
	printDesc(5)
	fmt.Println("print descreasing increasing")
	printDescIncreasing(5)
}

// input: 5
// output: 1 2 3 4 5
func printAsc(n int) {
	if n == 0 {
		return
	}
	printAsc(n - 1)
	fmt.Println(n)
}

func printDesc(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printDesc(n - 1)
}

func printDescIncreasing(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printDescIncreasing(n - 1)
	fmt.Println(n)

}
