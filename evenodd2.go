package main

import "fmt"

func main(){
	var num = 0
	fmt.Println("Enter a number")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println("Expecting number")
	} else {
		if isEven(num) {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

func isEven(num int) bool{
	if num % 2 == 0{
		return true
	}
	return false
}
