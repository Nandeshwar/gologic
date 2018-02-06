package main

import "fmt"

func main(){
	var begin = 0
	var end = 0
	fmt.Println("Enter start and end number")
	_, err := fmt.Scanf("%d %d", &begin, &end)

	if err != nil {
		fmt.Println("Expecting number")
	} else {
		evenList, oddList := evenOddList(begin, end)
		fmt.Println("Even number list = ",  evenList)
		fmt.Println("Odd Number list = ",  oddList)
	}
}

func evenOddList(begin int, end int) (evenList , oddList [] int) {
	for num := begin; num <= end; num++ {
		if isEven(num){
			evenList = append(evenList, num)
		} else {
			oddList = append(oddList, num)
		}
	}
	return
}

func isEven(num int) bool{
	if num % 2 == 0{
		return true
	}
	return false
}