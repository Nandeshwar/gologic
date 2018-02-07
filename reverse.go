package main

import "fmt"

func main() {
	var num = 0
	var reverse = 0
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &num)

	for num > 0 {
		rem := num % 10
		num = num / 10
		reverse = reverse * 10 + rem
	}

	fmt.Println(reverse)
}
