package main

import (
	"fmt"
)

/*
input: 5
output:
    [1]
    [1 1]
    [1 2 1]
    [1 3 3 1
    [1 4 6 4 1]


*/
func main() {
	fmt.Println(pascalTriangle(5))
}

func pascalTriangle(n int) [][]int {
	var outerArr [][]int
	outerArr = append(outerArr, []int{1})

	for i := 1; i < n; i++ {
		newArr := []int{1}
		prevArr := outerArr[i-1]
		for j := 1; j < i; j++ {
			sum := prevArr[j-1] + prevArr[j]
			newArr = append(newArr, sum)
		}
		newArr = append(newArr, 1)
		outerArr = append(outerArr, newArr)
	}
	return outerArr
}
