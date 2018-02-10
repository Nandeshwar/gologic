package main

import "fmt"

func main(){
	num := [][] int {{1, 2},
	                  {3, 4}}
	fmt.Println(transposeMatrix(num))
}

func transposeMatrix(num [][]int)(transpose [][]int){
	transpose = make([][]int, 2)
	for i := 0; i < 2; i++{
		transpose[i] = make([]int, 2)
		for j :=0; j <2; j++{
			transpose[i][j] = num[j][i]
		}
	}
	return transpose
}